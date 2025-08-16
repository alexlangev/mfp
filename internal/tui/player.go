package tui

import (
	"fmt"

	"github.com/alexlangev/mfp/internal/episodes"
	tea "github.com/charmbracelet/bubbletea"
)

type PModel struct {
	ep episodes.Episode
}

func (m PModel) Init() tea.Cmd {
	return nil
}

func (m PModel) Update(msg tea.Msg) (PModel, tea.Cmd) {
	// return to list
	// play / stop

	switch msg := msg.(type) {

	case SelectedMsg:
		m.ep = msg.selected

	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			return m, switchViewCmd(viewList)
		}
	}

	return m, nil
}

func (m PModel) View() string {
	s := fmt.Sprintf("Player view:\nlistening to: %s", m.ep.Title)

	return s
}

func NewPlayerView() PModel {
	return PModel{}
}
