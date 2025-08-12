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

		// TODO switch to player view
		return m, pCmd

	case tea.KeyMsg:
		switch msg.String() {

		// Move these to the subviews?
		case "ctrl+c", "q":
			return m, tea.Quit

		case "c":
			return m.switchView(viewConnecting)

		case "l":
			return m.switchView(viewList)

		case "p":
			return m.switchView(viewPlayer)
		}

	case EpisodesMsg:
		var cCmd, lCmd tea.Cmd
		m.connectingView, cCmd = m.connectingView.Update(msg)
		m.epList, lCmd = m.epList.Update(msg)
		// switch to list view?
		return m, tea.Batch(cCmd, lCmd)
	}

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

func (m model) switchView(target viewState) (model, tea.Cmd) {
	m.viewState = target

	if !m.inits[target] {
		m.inits[target] = true

		switch target {
		case viewConnecting:
			return m, m.connectingView.Init()
		case viewList:
			return m, m.epList.Init()
		case viewPlayer:
			return m, m.player.Init()
		}
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
