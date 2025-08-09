package tui

import tea "github.com/charmbracelet/bubbletea"

type EpModel struct{}

func (m EpModel) Init() tea.Cmd {
	return nil
}

func (m EpModel) Update(msg tea.Msg) (EpModel, tea.Cmd) {
	// No behavior yet
	return m, nil
}

func (m EpModel) View() string {
	return "List view\n"
}

func NewListView() EpModel {
	return EpModel{}
}
