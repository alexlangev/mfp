package tui

import (
	"fmt"

	"github.com/alexlangev/mfp/internal/episodes"
	tea "github.com/charmbracelet/bubbletea"
)

type EpModel struct {
	episodesLoaded bool
	epsList        episodes.Episodes
}

func (m EpModel) Init() tea.Cmd {
	return nil
}

func (m EpModel) Update(msg tea.Msg) (EpModel, tea.Cmd) {
	switch msg := msg.(type) {
	case EpisodesMsg:
		m.episodesLoaded = true
		m.epsList = msg.eps

		return m, nil
	}

	return m, nil
}

func (m EpModel) View() string {
	s := ""

	if m.episodesLoaded {
		for _, e := range m.epsList {
			s += fmt.Sprintf("%s\n\n", e.Title)
		}

	} else {
		s += "no episodes yet..."
	}
	return s
}

func NewListView() EpModel {
	return EpModel{}
}
