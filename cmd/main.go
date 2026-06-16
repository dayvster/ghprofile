package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"ghprofile/config"
	"ghprofile/github"
	"ghprofile/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	flag.Usage = func() {
		fmt.Println(`ghprofile: Pretty GitHub profile viewer

Usage:
	ghprofile [flags]

Flags:
	-u, --user        GitHub username(s) to fetch (comma-separated for multiple)
	-n                How many top repos to show (default: 5)
	--token           GitHub personal access token (or GITHUB_TOKEN env)
	--config          Config file path (default: ~/.config/ghprofile/config.json)
	--json            Output as JSON
	--format          Output format: text, json (default: text)
	--no-icons        Disable icons in the output
	--no-border       Remove card border from output
	--no-style        Remove all styles from output
	--no-demo         Do not fall back to demo data on fetch error; exit instead
	--demo            Force demo data (skip network and cache)
	--size            Output size: small, medium, large, full (default: medium)
	--tui             Use interactive TUI mode (default: true)
	-h, --help        Show this help message`)
	}
	userLong := flag.String("user", "", "GitHub username(s) to fetch (comma-separated)")
	userShort := flag.String("u", "", "GitHub username (shorthand)")
	topN := flag.Int("n", 5, "How many top repos to show")
	tokenFlag := flag.String("token", "", "GitHub personal access token (or GITHUB_TOKEN env)")
	configFlag := flag.String("config", "", "Config file path (default: ~/.config/ghprofile/config.json)")
	jsonFlag := flag.Bool("json", false, "Output as JSON")
	format := flag.String("format", "text", "Output format: text, json")
	noIcons := flag.Bool("no-icons", false, "Disable icons in the output")
	noDemo := flag.Bool("no-demo", false, "Do not fall back to demo data on fetch error; exit instead")
	demo := flag.Bool("demo", false, "Force demo data (skip network and cache)")
	noBorder := flag.Bool("no-border", false, "Remove card border from output")
	noStyle := flag.Bool("no-style", false, "Remove all styles from output")
	size := flag.String("size", "medium", "Output size: small, medium, large, full")
	tui := flag.Bool("tui", true, "Use interactive TUI mode")
	flag.Parse()

	user := *userLong
	if *userShort != "" {
		user = *userShort
	}

	// If no user provided and not running in demo mode, require --user
	if user == "" && !*demo {
		fmt.Fprintln(os.Stderr, "error: --user is required unless --demo is set")
		flag.Usage()
		os.Exit(2)
	}

	token := *tokenFlag
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	if token == "" {
		cfgPath := *configFlag
		if cfgPath == "" {
			cfgPath, _ = config.Path()
		}
		if cfgPath != "" {
			if cfg, err := config.Load(cfgPath); err == nil && cfg.Token != "" {
				token = cfg.Token
			}
		}
	}

	gh := &github.Github{Client: http.DefaultClient}
	if token != "" {
		gh.Token = token
	}

	useJSON := *jsonFlag || *format == "json"

	// Parse comma-separated usernames
	users := strings.Split(user, ",")
	for i := range users {
		users[i] = strings.TrimSpace(users[i])
	}

	render := func(p *github.Profile, repos []github.Repo) {
		if *tui {
			tea.NewProgram(ui.NewModel(p, repos, *topN, !*noIcons, *noBorder, *noStyle, *size)).Run()
		} else {
			ui.PrintProfile(p, repos, *topN, !*noIcons, *noBorder, *noStyle, *size)
		}
	}

	fetchUser := func(username string) (*github.Profile, []github.Repo) {
		if *demo {
			p, repos := github.DemoProfile(github.DemoProfileConfig{Username: username})
			return p, repos
		}

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		p, repos, err := gh.FetchProfileWithRepos(ctx, username)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: fetch failed for %s: %v\n", username, err)
			if *noDemo {
				if cp, cr, cerr := github.TryLoadCache(username); cerr == nil {
					fmt.Fprintf(os.Stderr, "loaded cached profile for %s\n", username)
					p, repos = cp, cr
				} else {
					fmt.Fprintf(os.Stderr, "no cache available and --no-demo set; exiting\n")
					os.Exit(1)
				}
			} else {
				if cp, cr, cerr := github.TryLoadCache(username); cerr == nil {
					fmt.Fprintf(os.Stderr, "warning: fetch failed — using cached data for %s\n", username)
					p, repos = cp, cr
				} else {
					fmt.Fprintf(os.Stderr, "warning: fetch failed (%v) — falling back to demo data for %s\n", err, username)
					p, repos = github.DemoProfile(github.DemoProfileConfig{Username: username})
				}
			}
		} else {
			if err := github.SaveCache(username, p, repos); err != nil {
				fmt.Fprintf(os.Stderr, "warning: failed to save cache: %v\n", err)
			}
		}
		return p, repos
	}

	if useJSON {
		results := make([]map[string]interface{}, 0, len(users))
		for _, u := range users {
			p, repos := fetchUser(u)
			results = append(results, map[string]interface{}{
				"username": u,
				"profile":  p,
				"repos":    repos,
			})
		}
		if len(results) == 1 {
			json.NewEncoder(os.Stdout).Encode(results[0])
		} else {
			json.NewEncoder(os.Stdout).Encode(results)
		}
		return
	}

	for i, u := range users {
		if i > 0 {
			fmt.Println(strings.Repeat("─", 60))
		}
		p, repos := fetchUser(u)
		render(p, repos)
	}
}
