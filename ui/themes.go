package ui

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Name         string
	BrightFg     lipgloss.Color
	AccentCyan   lipgloss.Color
	AccentBlue   lipgloss.Color
	AccentGreen  lipgloss.Color
	AccentOrange lipgloss.Color
	AccentPurple lipgloss.Color
	Subtle       lipgloss.Color
	BorderFg     lipgloss.Color
}

var Themes = map[string]Theme{
	"nord": {
		Name:         "nord",
		BrightFg:     lipgloss.Color("#eceff4"),
		AccentCyan:   lipgloss.Color("#88c0d0"),
		AccentBlue:   lipgloss.Color("#81a1c1"),
		AccentGreen:  lipgloss.Color("#a3be8c"),
		AccentOrange: lipgloss.Color("#d08770"),
		AccentPurple: lipgloss.Color("#b48ead"),
		Subtle:       lipgloss.Color("#4c566a"),
		BorderFg:     lipgloss.Color("#81a1c1"),
	},
	"dracula": {
		Name:         "dracula",
		BrightFg:     lipgloss.Color("#f8f8f2"),
		AccentCyan:   lipgloss.Color("#8be9fd"),
		AccentBlue:   lipgloss.Color("#6272a4"),
		AccentGreen:  lipgloss.Color("#50fa7b"),
		AccentOrange: lipgloss.Color("#ffb86c"),
		AccentPurple: lipgloss.Color("#bd93f9"),
		Subtle:       lipgloss.Color("#44475a"),
		BorderFg:     lipgloss.Color("#6272a4"),
	},
	"tokyonight": {
		Name:         "tokyonight",
		BrightFg:     lipgloss.Color("#c0caf5"),
		AccentCyan:   lipgloss.Color("#7dcfff"),
		AccentBlue:   lipgloss.Color("#7aa2f7"),
		AccentGreen:  lipgloss.Color("#9ece6a"),
		AccentOrange: lipgloss.Color("#ff9e64"),
		AccentPurple: lipgloss.Color("#bb9af7"),
		Subtle:       lipgloss.Color("#3b4261"),
		BorderFg:     lipgloss.Color("#7aa2f7"),
	},
	"monokai": {
		Name:         "monokai",
		BrightFg:     lipgloss.Color("#f8f8f2"),
		AccentCyan:   lipgloss.Color("#66d9ef"),
		AccentBlue:   lipgloss.Color("#66d9ef"),
		AccentGreen:  lipgloss.Color("#a1efe4"),
		AccentOrange: lipgloss.Color("#fd971f"),
		AccentPurple: lipgloss.Color("#ae81ff"),
		Subtle:       lipgloss.Color("#75715e"),
		BorderFg:     lipgloss.Color("#66d9ef"),
	},
	"gruvbox": {
		Name:         "gruvbox",
		BrightFg:     lipgloss.Color("#ebdbb2"),
		AccentCyan:   lipgloss.Color("#8ec07c"),
		AccentBlue:   lipgloss.Color("#83a598"),
		AccentGreen:  lipgloss.Color("#b8bb26"),
		AccentOrange: lipgloss.Color("#fe8019"),
		AccentPurple: lipgloss.Color("#d3869b"),
		Subtle:       lipgloss.Color("#928374"),
		BorderFg:     lipgloss.Color("#83a598"),
	},
	"solarized": {
		Name:         "solarized",
		BrightFg:     lipgloss.Color("#fdf6e3"),
		AccentCyan:   lipgloss.Color("#2aa198"),
		AccentBlue:   lipgloss.Color("#268bd2"),
		AccentGreen:  lipgloss.Color("#859900"),
		AccentOrange: lipgloss.Color("#cb4b16"),
		AccentPurple: lipgloss.Color("#6c71c4"),
		Subtle:       lipgloss.Color("#93a1a1"),
		BorderFg:     lipgloss.Color("#268bd2"),
	},
	"catppuccin": {
		Name:         "catppuccin",
		BrightFg:     lipgloss.Color("#cdd6f4"),
		AccentCyan:   lipgloss.Color("#74c7ec"),
		AccentBlue:   lipgloss.Color("#89b4fa"),
		AccentGreen:  lipgloss.Color("#a6e3a1"),
		AccentOrange: lipgloss.Color("#fab387"),
		AccentPurple: lipgloss.Color("#cba6f7"),
		Subtle:       lipgloss.Color("#585b70"),
		BorderFg:     lipgloss.Color("#89b4fa"),
	},
	"onelight": {
		Name:         "onelight",
		BrightFg:     lipgloss.Color("#383a42"),
		AccentCyan:   lipgloss.Color("#0184bc"),
		AccentBlue:   lipgloss.Color("#4078f2"),
		AccentGreen:  lipgloss.Color("#50a14f"),
		AccentOrange: lipgloss.Color("#d19a66"),
		AccentPurple: lipgloss.Color("#a626a4"),
		Subtle:       lipgloss.Color("#a0a1a7"),
		BorderFg:     lipgloss.Color("#4078f2"),
	},
}

var CurrentTheme = Themes["tokyonight"]

func SetTheme(name string) bool {
	if theme, ok := Themes[name]; ok {
		CurrentTheme = theme
		return true
	}
	return false
}

func GetAvailableThemes() []string {
	themes := make([]string, 0, len(Themes))
	for name := range Themes {
		themes = append(themes, name)
	}
	return themes
}
