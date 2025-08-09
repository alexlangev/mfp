package tui

import tea "github.com/charmbracelet/bubbletea"

type PModel struct{}

func (m PModel) Init() tea.Cmd {
	return nil
}

func (m PModel) Update(msg tea.Msg) (PModel, tea.Cmd) {
	// No behavior yet
	return m, nil
}

func (m PModel) View() string {
	return "Player view\n"
}

func NewPlayerView() PModel {
	return PModel{}
}
