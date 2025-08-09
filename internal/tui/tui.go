package tui

import (
	"github.com/alexlangev/mfp/internal/tui/connecting"
	"github.com/alexlangev/mfp/internal/tui/episodeList"
	"github.com/alexlangev/mfp/internal/tui/player"
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
	// subviews
	connectingView connecting.Model
	epList         episodeList.Model
	player         player.Model
}

func (m model) Init() tea.Cmd {

	m.inits = map[viewState]bool{
		viewConnecting: true,
		viewList:       false,
		viewPlayer:     false,
	}

	return m.connectingView.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "c":
			return m.switchView(viewConnecting)

		case "l":
			return m.switchView(viewList)

		case "p":
			return m.switchView(viewPlayer)
		}
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
		connectingView: connecting.NewConnectingView(),
		epList:         episodeList.NewListView(),
		player:         player.NewPlayerView(),
		inits: map[viewState]bool{
			viewConnecting: false,
			viewList:       false,
			viewPlayer:     false,
		},
	}
}
