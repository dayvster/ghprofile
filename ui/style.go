package ui

import "github.com/charmbracelet/lipgloss"

// Panel returns a lipgloss.Style configured for the card panel.
// width: if 0 then no width constraint (full width), otherwise sets Width(width).
func Panel(width int) lipgloss.Style {
	s := lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(CurrentTheme.BorderFg).
		Foreground(CurrentTheme.BrightFg)
	if width > 0 {
		s = s.Width(width)
	}
	return s
}

// GetStyles returns theme-aware styles. Call this after SetTheme.
func GetStyles() map[string]lipgloss.Style {
	return map[string]lipgloss.Style{
		"title": lipgloss.NewStyle().
			Bold(true).
			Foreground(CurrentTheme.AccentCyan).
			MarginBottom(0),

		"url": lipgloss.NewStyle().
			Foreground(CurrentTheme.AccentBlue).
			Underline(true),

		"subtle": lipgloss.NewStyle().
			Foreground(CurrentTheme.Subtle),

		"accent": lipgloss.NewStyle().
			Bold(true).
			Foreground(CurrentTheme.AccentPurple),

		"icon": lipgloss.NewStyle().
			Bold(true).
			Foreground(CurrentTheme.AccentCyan),

		"repo_title": lipgloss.NewStyle().
			Bold(true).
			Foreground(CurrentTheme.AccentOrange),

		"value": lipgloss.NewStyle().
			Bold(true).
			Foreground(CurrentTheme.AccentGreen),
	}
}

var (
	TitleStyle lipgloss.Style
	URLStyle   lipgloss.Style
	Subtle     lipgloss.Style
	Accent     lipgloss.Style
	IconStyle  lipgloss.Style
	RepoTitle  lipgloss.Style
	ValueStyle lipgloss.Style
)

func init() {
	UpdateStyles()
}

func UpdateStyles() {
	styles := GetStyles()
	TitleStyle = styles["title"]
	URLStyle = styles["url"]
	Subtle = styles["subtle"]
	Accent = styles["accent"]
	IconStyle = styles["icon"]
	RepoTitle = styles["repo_title"]
	ValueStyle = styles["value"]
}
