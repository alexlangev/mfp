package tui

import (
	"github.com/alexlangev/mfp/internal/episodes"
	tea "github.com/charmbracelet/bubbletea"
)

type viewState int

const (
	viewConnecting viewState = iota
	viewList
	viewPlayer
)

type model struct {
	viewState viewState
	inits     map[viewState]bool
	eps       episodes.Episodes
	selected  episodes.Episode
	// subviews
	// TODO find better naming?
	connectingView ConModel
	epList         EpModel
	player         PModel
}

type EpisodesMsg struct {
	eps episodes.Episodes
}

type SelectedMsg struct {
	selected episodes.Episode
}

type SwitchViewMsg struct {
	viewState viewState
}

func (m model) Init() tea.Cmd {
	m.inits = map[viewState]bool{
		viewConnecting: true,
		viewList:       false,
		viewPlayer:     false,
	}

	// fetch and parse episode
	return tea.Batch(
		m.connectingView.Init(),
		fetchEpisodesCmd(),
	)
}

func fetchEpisodesCmd() tea.Cmd {
	return func() tea.Msg {
		episodes, _ := episodes.GetEpisodes()
		return EpisodesMsg{eps: episodes}
	}
}

func switchViewCmd(v viewState) tea.Cmd {
	return func() tea.Msg {
		return SwitchViewMsg{viewState: v}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		var cCmd, lCmd, pCmd tea.Cmd
		m.connectingView, cCmd = m.connectingView.Update(msg)
		m.epList, lCmd = m.epList.Update(msg)
		m.player, pCmd = m.player.Update(msg)

		return m, tea.Batch(cCmd, lCmd, pCmd)

	case SelectedMsg:
		m.selected = msg.selected
		var pCmd tea.Cmd
		m.player, pCmd = m.player.Update(msg)

		return m, pCmd

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit
		}

	case SwitchViewMsg:
		m.viewState = msg.viewState
		var cmd tea.Cmd

		if !m.inits[msg.viewState] {
			m.inits[msg.viewState] = true

			switch msg.viewState {
			case viewConnecting:
				cmd = m.connectingView.Init()
			case viewList:
				cmd = m.epList.Init()
			case viewPlayer:
				cmd = m.player.Init()
			}
		}
		return m, cmd

	case EpisodesMsg:
		var cCmd, lCmd, pCmd tea.Cmd
		m.connectingView, cCmd = m.connectingView.Update(msg)
		m.epList, lCmd = m.epList.Update(msg)

		// not sure about this
		pCmd = switchViewCmd(viewList)

		return m, tea.Batch(cCmd, lCmd, pCmd)
	}

	// msg handling of subviews
	switch m.viewState {
	case viewConnecting:
		var cmd tea.Cmd
		m.connectingView, cmd = m.connectingView.Update(msg)
		return m, cmd

	case viewList:
		var cmd tea.Cmd
		m.epList, cmd = m.epList.Update(msg)
		return m, cmd

	case viewPlayer:
		var cmd tea.Cmd
		m.player, cmd = m.player.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch m.viewState {
	case viewConnecting:
		return m.connectingView.View()
	case viewList:
		return m.epList.View()
	case viewPlayer:
		return m.player.View()
	}
	return ""
}

func InitialModel() model {
	return model{
		viewState:      viewConnecting,
		connectingView: NewConnectingView(),
		epList:         NewListView(),
		player:         NewPlayerView(),
		inits: map[viewState]bool{
			viewConnecting: false,
			viewList:       false,
			viewPlayer:     false,
		},
	}
}
