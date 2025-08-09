package tui

import tea "github.com/charmbracelet/bubbletea"

type playerView struct{}

func (p playerView) Init() tea.Cmd {
	return nil
}

func (p playerView) Update(msg tea.Msg) (playerView, tea.Cmd) {
	// No behavior yet
	return p, nil
}

func (p playerView) View() string {
	return "Player view\n"
}

func NewPlayerView() playerView {
	return playerView{}
}
