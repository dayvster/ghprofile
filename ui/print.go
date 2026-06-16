package ui

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"ghprofile/github"

	"github.com/charmbracelet/lipgloss"
)

func PrintProfile(p *github.Profile, repos []github.Repo, topN int, showIcons bool, noBorder bool, noStyle bool, size string) {
	fmt.Print(FormatProfile(p, repos, topN, showIcons, noBorder, noStyle, size))
}

func FormatProfile(p *github.Profile, repos []github.Repo, topN int, showIcons bool, noBorder bool, noStyle bool, size string) string {
	if p == nil {
		return "No profile\n"
	}

	var b strings.Builder
	title := TitleStyle.Render(fmt.Sprintf("%s", p.FullName))
	url := URLStyle.Render(p.URL)
	iconRender := func(s string) string {
		if !showIcons || s == "" {
			return ""
		}
		return IconStyle.Render(s)
	}
	if iconRender(IconUser) != "" {
		b.WriteString(iconRender(IconUser) + "  " + title + "  " + url + "\n")
	} else {
		b.WriteString(title + "  " + url + "\n")
	}
	if p.Bio != "" {
		b.WriteString(Subtle.Render(p.Bio) + "\n")
	}
	b.WriteString("\n")
	type statEntry struct{ icon, label, value string }
	totalStars := p.TotalStars
	totalForks := p.TotalForks
	avg := p.AvgStarsPerRepo

	stats := []statEntry{
		{iconRender(IconFollowers), "Followers:", fmt.Sprintf("%d", p.FollowersAmount)},
		{iconRender(IconFollowing), "Following:", fmt.Sprintf("%d", p.FollowingAmount)},
		{iconRender(IconRepo), "Public repos:", fmt.Sprintf("%d", p.PublicReposAmount)},
		{iconRender(IconGist), "Public gists:", fmt.Sprintf("%d", p.PublicGistsAmount)},
		{iconRender(IconStar), "Total stars:", fmt.Sprintf("%d", totalStars)},
		{iconRender(IconFork), "Total forks:", fmt.Sprintf("%d", totalForks)},
		{iconRender(IconStar), "Avg stars/repo:", fmt.Sprintf("%.2f", avg)},
	}
	maxLabel := 0
	for _, s := range stats {
		if len(s.label) > maxLabel {
			maxLabel = len(s.label)
		}
	}
	for _, s := range stats {
		if s.icon != "" {
			b.WriteString(s.icon + "  ")
		} else {
			b.WriteString("   ")
		}
		padded := fmt.Sprintf("%-*s", maxLabel, s.label)
		b.WriteString(Accent.Render(padded) + " " + ValueStyle.Render(s.value) + "\n")
	}

	if p.ContributionCalendar != nil {
		b.WriteString(renderContributions(p.ContributionCalendar, iconRender))
	}

	type langEntry struct {
		name  string
		bytes int
		count int
	}

	var langs []langEntry
	hasByteData := false

	for _, r := range repos {
		if r.Languages != nil {
			hasByteData = true
			break
		}
	}

	if hasByteData {
		langBytes := map[string]int{}
		for _, r := range repos {
			for lang, bytes := range r.Languages {
				langBytes[lang] += bytes
			}
		}
		totalBytes := 0
		for _, bytes := range langBytes {
			totalBytes += bytes
		}
		for name, bytes := range langBytes {
			langs = append(langs, langEntry{name: name, bytes: bytes})
		}
		sort.Slice(langs, func(i, j int) bool { return langs[i].bytes > langs[j].bytes })

		if totalBytes > 0 && len(langs) > 0 {
			b.WriteString("\n")
			b.WriteString(Subtle.Render("Languages:") + "\n")
			for _, x := range langs {
				pct := float64(x.bytes) / float64(totalBytes) * 100
				icon := GetLangIcon(x.name)
				out := fmt.Sprintf("%s  %s: %.1f%%\n", iconRender(icon), Accent.Render(x.name), pct)
				if icon == "" {
					out = fmt.Sprintf("   %s: %.1f%%\n", Accent.Render(x.name), pct)
				}
				b.WriteString(out)
			}
		}
	} else {
		langCount := map[string]int{}
		for _, r := range repos {
			lang := r.Language
			if lang == "" {
				continue
			}
			langCount[lang]++
		}
		for name, count := range langCount {
			langs = append(langs, langEntry{name: name, count: count})
		}
		sort.Slice(langs, func(i, j int) bool { return langs[i].count > langs[j].count })

		if len(langs) > 0 {
			b.WriteString("\n")
			b.WriteString(Subtle.Render("Languages:") + "\n")
			for _, x := range langs {
				icon := GetLangIcon(x.name)
				out := fmt.Sprintf("%s  %s: %d\n", iconRender(icon), Accent.Render(x.name), x.count)
				if icon == "" {
					out = fmt.Sprintf("   %s: %d\n", Accent.Render(x.name), x.count)
				}
				b.WriteString(out)
			}
		}
	}

	if topN > len(repos) {
		topN = len(repos)
	}
	if topN > 0 {
		b.WriteString("\n")
		if noStyle {
			b.WriteString("Top repos:\n")
		} else {
			b.WriteString(Subtle.Render("Top repos:") + "\n")
		}
		for i := 0; i < topN; i++ {
			r := repos[i]
			langIcon := GetLangIcon(r.Language)
			if noStyle {
				b.WriteString(fmt.Sprintf("%d. %s %s ★ %d  %s %d\n", i+1, r.FullName, r.Language, r.StargazersCount, IconFork, r.ForksCount))
				b.WriteString("  " + r.HTMLURL + "\n")
			} else {
				b.WriteString(fmt.Sprintf("%d. %s %s %s %d  %s %d\n", i+1, RepoTitle.Render(r.FullName), iconRender(langIcon), iconRender("★"), r.StargazersCount, iconRender(IconFork), r.ForksCount))
				b.WriteString("  " + URLStyle.Render(r.HTMLURL) + "\n")
			}
		}
	}

	out := b.String()
	if noStyle || noBorder {
		return out
	}

	width := 0
	switch strings.ToLower(size) {
	case "small":
		width = 48
	case "medium":
		width = 88
	case "large":
		width = 120
	case "full":
		width = 0
	default:
		width = 88
	}

	panel := Panel(width).Render(out)
	return lipgloss.NewStyle().Margin(1, 2).Render(panel) + "\n"
}

func renderContributions(cal *github.ContributionCalendar, iconRender func(string) string) string {
	if cal == nil || len(cal.Weeks) == 0 {
		return ""
	}
	var buf strings.Builder
	buf.WriteString("\n")
	buf.WriteString(Subtle.Render("Contributions:") + "\n")
	buf.WriteString(fmt.Sprintf("%s %s %d\n", iconRender(IconStar), Accent.Render("Total:"), cal.TotalContributions))

	weekly := make([]int, len(cal.Weeks))
	maxW := 0
	for i, w := range cal.Weeks {
		s := 0
		for _, d := range w.Days {
			s += d.Count
		}
		weekly[i] = s
		if s > maxW {
			maxW = s
		}
	}

	if maxW > 0 {
		chars := []rune("▁▂▃▄▅▆▇█")
		for _, s := range weekly {
			l := int(float64(s) / float64(maxW) * 7)
			if l > 7 {
				l = 7
			}
			buf.WriteRune(chars[l])
		}
		buf.WriteString("\n")
	}

	return buf.String()
}

func RenderJSON(p *github.Profile, repos []github.Repo) string {
	data := struct {
		Profile *github.Profile `json:"profile"`
		Repos   []github.Repo   `json:"repos"`
	}{p, repos}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error":%q}`, err)
	}
	return string(b)
}
