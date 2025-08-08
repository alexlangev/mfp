package tui

import tea "github.com/charmbracelet/bubbletea"

type ViewState int

const (
	ViewConnecting ViewState = iota
	ViewList
	ViewPlayer
)

type model struct {
	viewState ViewState
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "c":
			m.viewState = ViewConnecting

		case "l":
			m.viewState = ViewList

		case "p":
			m.viewState = ViewPlayer
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "current view is: "

	switch m.viewState {
	case ViewConnecting:
		s += "Connecting"
	case ViewList:
		s += "List"
	case ViewPlayer:
		s += "Player"
	}

	s = s + "\n"
	return s
}

func InitialModel() model {
	return model{
		viewState: ViewConnecting,
	}
}
