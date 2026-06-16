package ui

import (
	"ghprofile/github"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	profile   *github.Profile
	repos     []github.Repo
	topN      int
	showIcons bool
	noBorder  bool
	noStyle   bool
	size      string
}

type renderMsg struct{}

func NewModel(p *github.Profile, repos []github.Repo, topN int, showIcons, noBorder, noStyle bool, size string) tea.Model {
	return &model{
		profile:   p,
		repos:     repos,
		topN:      topN,
		showIcons: showIcons,
		noBorder:  noBorder,
		noStyle:   noStyle,
		size:      size,
	}
}

func (m *model) Init() tea.Cmd {
	return func() tea.Msg {
		return renderMsg{}
	}
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m *model) View() string {
	return FormatProfile(m.profile, m.repos, m.topN, m.showIcons, m.noBorder, m.noStyle, m.size)
}
