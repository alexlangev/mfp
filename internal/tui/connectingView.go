package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type ConModel struct {
	isConnecting bool
	spinner      spinner.Model
}

type connectedMsg struct{}

func (m ConModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
	)
}

func (m ConModel) Update(msg tea.Msg) (ConModel, tea.Cmd) {
	switch msg := msg.(type) {
	case EpisodesMsg:
		m.isConnecting = false
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m ConModel) View() string {
	s := ""

	if m.isConnecting {
		s = fmt.Sprintf("Trying to connect%s", m.spinner.View())
	} else {
		s = "Connection established!!!"
	}

	return s
}

func NewConnectingView() ConModel {
	sp := spinner.New()
	sp.Spinner = spinner.Ellipsis

	return ConModel{
		isConnecting: true,
		spinner:      sp,
	}
}
