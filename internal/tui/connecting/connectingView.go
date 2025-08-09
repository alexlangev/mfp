package connecting

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	isConnecting bool
	spinner      spinner.Model
}

type connectedMsg struct{}

func (m Model) Init() tea.Cmd {
	// return m.spinner.Tick
	return tea.Batch(
		m.spinner.Tick,
		mockConnection(),
	)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case connectedMsg:
		m.isConnecting = false
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	s := ""

	if m.isConnecting {
		s = fmt.Sprintf("Trying to connect%s", m.spinner.View())
	} else {
		s = "Connection established!"
	}

	return s
}

func NewConnectingView() Model {
	sp := spinner.New()
	sp.Spinner = spinner.Ellipsis

	return Model{
		isConnecting: true,
		spinner:      sp,
	}
}

func mockConnection() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(15 * time.Second)
		return connectedMsg{}
	}
}
