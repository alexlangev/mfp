package tui

import (
	"github.com/alexlangev/mfp/internal/episodes"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type EpModel struct {
	episodesLoaded bool
	epsList        episodes.Episodes
	list           list.Model
}

type epListItem struct {
	ep episodes.Episode
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func (li epListItem) Title() string       { return li.ep.Title }
func (li epListItem) Description() string { return li.ep.Title }
func (li epListItem) FilterValue() string { return li.ep.Title }

func (m EpModel) Init() tea.Cmd {
	return nil
}

func (m EpModel) Update(msg tea.Msg) (EpModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case EpisodesMsg:
		m.episodesLoaded = true
		m.epsList = msg.eps

		items := make([]list.Item, len(m.epsList))
		for i, ep := range m.epsList {
			items[i] = epListItem{ep}
		}

		m.list.SetItems(items)

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ":
			// return m, selectEpisodeCmd(m)

			return m, tea.Batch(
				selectEpisodeCmd(m),
				switchViewCmd(viewPlayer),
			)
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func selectEpisodeCmd(m EpModel) tea.Cmd {
	return func() tea.Msg {
		selected := m.list.SelectedItem()

		if selected != nil {
			if epItem, ok := selected.(epListItem); ok {
				return SelectedMsg{selected: epItem.ep}
			}
		}
		return nil
	}
}

func (m EpModel) View() string {
	if m.episodesLoaded {
		return docStyle.Render(m.list.View())
	}
	return "no episodes yet..."
}

func NewListView() EpModel {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false

	l := list.New([]list.Item{}, delegate, 0, 0)
	l.Title = "Music for programing"

	return EpModel{
		list: l,
	}
}
