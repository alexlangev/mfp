package tui

import (
	"github.com/alexlangev/mfp/internal/episodes"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type EpModel struct {
	episodesLoaded bool
	epsList        episodes.Episodes // rename to list items or something?
	list           list.Model
}

type epListItem struct {
	ep episodes.Episode
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func (li epListItem) Title() string       { return li.ep.Title }
func (li epListItem) Description() string { return li.ep.Title } // won't use it, maybe...
func (li epListItem) FilterValue() string { return li.ep.Title }

func (m EpModel) Init() tea.Cmd {
	return nil
}

func (m EpModel) Update(msg tea.Msg) (EpModel, tea.Cmd) {
	switch msg := msg.(type) {
	case EpisodesMsg:
		m.episodesLoaded = true
		m.epsList = msg.eps

		items := make([]list.Item, len(m.epsList))
		for i, ep := range m.epsList {
			items[i] = epListItem{ep}
		}

		m.list.SetItems(items)

		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m EpModel) View() string {
	if m.episodesLoaded {
		return docStyle.Render(m.list.View())
	}
	return "no episodes yet..."
}

func NewListView() EpModel {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 50, 100)
	l.Title = "Episodes"

	return EpModel{
		list: l,
	}
}
