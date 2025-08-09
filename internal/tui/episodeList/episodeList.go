package episodeList

import tea "github.com/charmbracelet/bubbletea"

type Model struct{}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	// No behavior yet
	return m, nil
}

func (m Model) View() string {
	return "List view\n"
}

func NewListView() Model {
	return Model{}
}
