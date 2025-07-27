package tui

import (
	"github.com/alexlangev/mfp/internal/episodes"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		eps, err := episodes.GetEpisodes()
		if err != nil {
			return errMsg{err: err}
		}
		return epsMsg{episodes: eps}
	}
}
