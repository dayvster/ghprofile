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
	"ember": {
		Name:         "ember",
		BrightFg:     lipgloss.Color("#e0e0e0"),
		AccentCyan:   lipgloss.Color("#71c7d4"),
		AccentBlue:   lipgloss.Color("#5fafef"),
		AccentGreen:  lipgloss.Color("#8fb573"),
		AccentOrange: lipgloss.Color("#eb7d3e"),
		AccentPurple: lipgloss.Color("#a882d1"),
		Subtle:       lipgloss.Color("#555555"),
		BorderFg:     lipgloss.Color("#5fafef"),
	},
	"ayu-dark": {
		Name:         "ayu-dark",
		BrightFg:     lipgloss.Color("#e6e1cf"),
		AccentCyan:   lipgloss.Color("#95e6cb"),
		AccentBlue:   lipgloss.Color("#6dcbfa"),
		AccentGreen:  lipgloss.Color("#a6e22e"),
		AccentOrange: lipgloss.Color("#ffb454"),
		AccentPurple: lipgloss.Color("#d9d0ff"),
		Subtle:       lipgloss.Color("#626a73"),
		BorderFg:     lipgloss.Color("#6dcbfa"),
	},
	"ayu-light": {
		Name:         "ayu-light",
		BrightFg:     lipgloss.Color("#3d3d3d"),
		AccentCyan:   lipgloss.Color("#0d9c9c"),
		AccentBlue:   lipgloss.Color("#2196f3"),
		AccentGreen:  lipgloss.Color("#86a02a"),
		AccentOrange: lipgloss.Color("#ff8f40"),
		AccentPurple: lipgloss.Color("#7253b4"),
		Subtle:       lipgloss.Color("#999999"),
		BorderFg:     lipgloss.Color("#2196f3"),
	},
	"one-dark": {
		Name:         "one-dark",
		BrightFg:     lipgloss.Color("#abb2bf"),
		AccentCyan:   lipgloss.Color("#56b6c2"),
		AccentBlue:   lipgloss.Color("#61afef"),
		AccentGreen:  lipgloss.Color("#98c379"),
		AccentOrange: lipgloss.Color("#e5c07b"),
		AccentPurple: lipgloss.Color("#c678dd"),
		Subtle:       lipgloss.Color("#5c6370"),
		BorderFg:     lipgloss.Color("#61afef"),
	},
	"material": {
		Name:         "material",
		BrightFg:     lipgloss.Color("#eeffff"),
		AccentCyan:   lipgloss.Color("#80cbc4"),
		AccentBlue:   lipgloss.Color("#82b1ff"),
		AccentGreen:  lipgloss.Color("#c3e88d"),
		AccentOrange: lipgloss.Color("#ffb74d"),
		AccentPurple: lipgloss.Color("#c792ea"),
		Subtle:       lipgloss.Color("#546e7a"),
		BorderFg:     lipgloss.Color("#82b1ff"),
	},
	"synthwave": {
		Name:         "synthwave",
		BrightFg:     lipgloss.Color("#f2f4f6"),
		AccentCyan:   lipgloss.Color("#00ff88"),
		AccentBlue:   lipgloss.Color("#00d9ff"),
		AccentGreen:  lipgloss.Color("#39ff14"),
		AccentOrange: lipgloss.Color("#ffee58"),
		AccentPurple: lipgloss.Color("#ff006e"),
		Subtle:       lipgloss.Color("#72849f"),
		BorderFg:     lipgloss.Color("#00d9ff"),
	},
	"palenight": {
		Name:         "palenight",
		BrightFg:     lipgloss.Color("#a6acb5"),
		AccentCyan:   lipgloss.Color("#89ddff"),
		AccentBlue:   lipgloss.Color("#82aaff"),
		AccentGreen:  lipgloss.Color("#c3e88d"),
		AccentOrange: lipgloss.Color("#ffb74d"),
		AccentPurple: lipgloss.Color("#c792ea"),
		Subtle:       lipgloss.Color("#546e7a"),
		BorderFg:     lipgloss.Color("#82aaff"),
	},
	"github-light": {
		Name:         "github-light",
		BrightFg:     lipgloss.Color("#24292f"),
		AccentCyan:   lipgloss.Color("#0969da"),
		AccentBlue:   lipgloss.Color("#0969da"),
		AccentGreen:  lipgloss.Color("#2da44e"),
		AccentOrange: lipgloss.Color("#d1540d"),
		AccentPurple: lipgloss.Color("#8250df"),
		Subtle:       lipgloss.Color("#57606a"),
		BorderFg:     lipgloss.Color("#0969da"),
	},
	"github-dark": {
		Name:         "github-dark",
		BrightFg:     lipgloss.Color("#c9d1d9"),
		AccentCyan:   lipgloss.Color("#58a6ff"),
		AccentBlue:   lipgloss.Color("#58a6ff"),
		AccentGreen:  lipgloss.Color("#3fb950"),
		AccentOrange: lipgloss.Color("#fb8500"),
		AccentPurple: lipgloss.Color("#d2a8ff"),
		Subtle:       lipgloss.Color("#6e7681"),
		BorderFg:     lipgloss.Color("#58a6ff"),
	},
	"nord-light": {
		Name:         "nord-light",
		BrightFg:     lipgloss.Color("#2e3440"),
		AccentCyan:   lipgloss.Color("#88c0d0"),
		AccentBlue:   lipgloss.Color("#81a1c1"),
		AccentGreen:  lipgloss.Color("#a3be8c"),
		AccentOrange: lipgloss.Color("#d08770"),
		AccentPurple: lipgloss.Color("#b48ead"),
		Subtle:       lipgloss.Color("#d8dee9"),
		BorderFg:     lipgloss.Color("#81a1c1"),
	},
	"everforest": {
		Name:         "everforest",
		BrightFg:     lipgloss.Color("#d4be98"),
		AccentCyan:   lipgloss.Color("#7fbbb3"),
		AccentBlue:   lipgloss.Color("#7daea3"),
		AccentGreen:  lipgloss.Color("#a7c957"),
		AccentOrange: lipgloss.Color("#e39937"),
		AccentPurple: lipgloss.Color("#d196c6"),
		Subtle:       lipgloss.Color("#523f38"),
		BorderFg:     lipgloss.Color("#7daea3"),
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
