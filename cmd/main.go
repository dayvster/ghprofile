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
	--theme           Color theme: nord, dracula, tokyonight, monokai, gruvbox, solarized, catppuccin, onelight
	--list-themes     List available themes
	--json            Output as JSON
	--format          Output format: text, json (default: text)
	--no-icons        Disable icons in the output
	--no-border       Remove card border from output
	--no-style        Remove all styles from output
	--no-demo         Do not fall back to demo data on fetch error; exit instead
	--demo            Force demo data (skip network and cache)
	--size            Output size: small, medium, large, full (default: medium)
	--tui             Use interactive TUI mode (default: true)
	--verbose         Print API request URLs, timing, cache info to stderr
	-v                Shorthand for --verbose
	--quiet           Suppress stderr warning messages
	-q                Shorthand for --quiet
	--sort            Sort repos by: stars, forks, name, updated (default: stars)
	--since           Only show repos updated after this date (YYYY-MM-DD)
	-h, --help        Show this help message`)
	}
	userLong := flag.String("user", "", "GitHub username(s) to fetch (comma-separated)")
	userShort := flag.String("u", "", "GitHub username (shorthand)")
	topN := flag.Int("n", 5, "How many top repos to show")
	tokenFlag := flag.String("token", "", "GitHub personal access token (or GITHUB_TOKEN env)")
	configFlag := flag.String("config", "", "Config file path (default: ~/.config/ghprofile/config.json)")
	themeFlag := flag.String("theme", "", "Color theme")
	listThemes := flag.Bool("list-themes", false, "List available themes")
	jsonFlag := flag.Bool("json", false, "Output as JSON")
	format := flag.String("format", "text", "Output format: text, json")
	noIcons := flag.Bool("no-icons", false, "Disable icons in the output")
	noDemo := flag.Bool("no-demo", false, "Do not fall back to demo data on fetch error; exit instead")
	demo := flag.Bool("demo", false, "Force demo data (skip network and cache)")
	noBorder := flag.Bool("no-border", false, "Remove card border from output")
	noStyle := flag.Bool("no-style", false, "Remove all styles from output")
	size := flag.String("size", "medium", "Output size: small, medium, large, full")
	tui := flag.Bool("tui", true, "Use interactive TUI mode")
	verboseLong := flag.Bool("verbose", false, "Print API request URLs, timing, cache info to stderr")
	verboseShort := flag.Bool("v", false, "Shorthand for --verbose")
	quietLong := flag.Bool("quiet", false, "Suppress stderr warning messages")
	quietShort := flag.Bool("q", false, "Shorthand for --quiet")
	sortBy := flag.String("sort", "stars", "Sort repos by: stars, forks, name, updated")
	since := flag.String("since", "", "Only show repos updated after this date (YYYY-MM-DD)")
	flag.Parse()

	if *listThemes {
		themes := ui.GetAvailableThemes()
		fmt.Println("Available themes:")
		for _, t := range themes {
			fmt.Printf("  • %s\n", t)
		}
		os.Exit(0)
	}

	isVerbose := *verboseLong || *verboseShort
	isQuiet := *quietLong || *quietShort

	user := *userLong
	if *userShort != "" {
		user = *userShort
	}

	if user == "" && !*demo {
		fmt.Fprintln(os.Stderr, "error: --user is required unless --demo is set")
		flag.Usage()
		os.Exit(2)
	}

	token := *tokenFlag
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}

	theme := *themeFlag
	if token == "" || theme == "" {
		cfgPath := *configFlag
		if cfgPath == "" {
			cfgPath, _ = config.Path()
		}
		if cfgPath != "" {
			if cfg, err := config.Load(cfgPath); err == nil {
				if token == "" && cfg.Token != "" {
					token = cfg.Token
				}
				if theme == "" && cfg.Theme != "" {
					theme = cfg.Theme
				}
			}
		}
	}

	gh := &github.Github{
		Client:  http.DefaultClient,
		Verbose: isVerbose,
		Quiet:   isQuiet,
	}
	if token != "" {
		gh.Token = token
	}

	useJSON := *jsonFlag || *format == "json"

	users := strings.Split(user, ",")
	for i := range users {
		users[i] = strings.TrimSpace(users[i])
	}

	render := func(p *github.Profile, repos []github.Repo) {
		if *tui {
			tea.NewProgram(ui.NewModel(p, repos, *topN, !*noIcons, *noBorder, *noStyle, *size, theme)).Run()
		} else {
			if theme != "" {
				ui.SetTheme(theme)
				ui.UpdateStyles()
			}
			ui.PrintProfile(p, repos, *topN, !*noIcons, *noBorder, *noStyle, *size)
		}
	}

	verbosePrint := func(format string, args ...interface{}) {
		if isVerbose {
			fmt.Fprintf(os.Stderr, format, args...)
		}
	}

	quietPrint := func(format string, args ...interface{}) {
		if !isQuiet {
			fmt.Fprintf(os.Stderr, format, args...)
		}
	}

	fetchUser := func(username string) (*github.Profile, []github.Repo) {
		if *demo {
			p, repos := github.DemoProfile(github.DemoProfileConfig{Username: username})
			return p, repos
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		cacheStart := time.Now()
		p, repos, err := gh.FetchProfileWithRepos(ctx, username)
		if err != nil {
			if cp, cr, cerr := github.TryLoadCache(username); cerr == nil {
				quietPrint("warning: fetch failed — using cached data for %s\n", username)
				verbosePrint("[verbose] cache hit for %s (%.0fs old)\n", username, time.Since(cacheStart).Seconds())
				p, repos = cp, cr
			} else if *noDemo {
				quietPrint("no cache available and --no-demo set; exiting\n")
				os.Exit(1)
			} else {
				quietPrint("warning: fetch failed (%v) — falling back to demo data for %s\n", err, username)
				p, repos = github.DemoProfile(github.DemoProfileConfig{Username: username})
			}
		} else {
			if err := github.SaveCache(username, p, repos); err != nil {
				verbosePrint("[verbose] cache write failed for %s: %v\n", username, err)
			} else {
				verbosePrint("[verbose] cache saved for %s\n", username)
			}
		}

		repos = github.FilterSince(repos, *since)
		github.SortRepos(repos, *sortBy)

		if !*demo {
			langCtx, langCancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer langCancel()
			verbosePrint("[verbose] fetching language breakdown for %d repos...\n", len(repos))
			gh.FetchLanguagesForRepos(langCtx, repos)
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
