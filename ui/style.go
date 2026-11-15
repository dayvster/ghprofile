package ui

import "github.com/charmbracelet/lipgloss"

var (
	brightFg     = lipgloss.Color("#e0e0e0")
	accentCyan   = lipgloss.Color("#5eadc1")
	accentBlue   = lipgloss.Color("#569cd6")
	accentGreen  = lipgloss.Color("#6a9955")
	accentOrange = lipgloss.Color("#daa520")
	accentPurple = lipgloss.Color("#c586c0")
)

// Panel returns a lipgloss.Style configured for the card panel.
// width: if 0 then no width constraint (full width), otherwise sets Width(width).
func Panel(width int) lipgloss.Style {
	s := lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(accentBlue).
		Foreground(brightFg)
	if width > 0 {
		s = s.Width(width)
	}
	return s
}

var (
	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(accentCyan).
		MarginBottom(0)

	URLStyle = lipgloss.NewStyle().
		Foreground(accentBlue).
		Underline(true)

	Subtle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888"))

	Accent = lipgloss.NewStyle().
		Bold(true).
		Foreground(accentPurple)

	IconStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(accentCyan)

	RepoTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(accentOrange)

	ValueStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(accentGreen)
)
