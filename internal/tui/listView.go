package tui

import tea "github.com/charmbracelet/bubbletea"

type listView struct{}

func (l listView) Init() tea.Cmd {
	return nil
}

func (l listView) Update(msg tea.Msg) (listView, tea.Cmd) {
	// No behavior yet
	return l, nil
}

func (l listView) View() string {
	return "List view\n"
}

func NewListView() listView {
	return listView{}
}
