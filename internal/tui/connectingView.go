package tui

import tea "github.com/charmbracelet/bubbletea"

type connectingView struct{}

func (c connectingView) Init() tea.Cmd {
	return nil
}

func (c connectingView) Update(msg tea.Msg) (connectingView, tea.Cmd) {
	return c, nil
}

func (c connectingView) View() string {
	return "Connecting view\n"
}

func NewConnectingView() connectingView {
	return connectingView{}
}
