package tui

import tea "github.com/charmbracelet/bubbletea"

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
	connectingView connectingView
	listView       listView
	playerView     playerView
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
		m.listView, cmd = m.listView.Update(msg)
		return m, cmd

	case viewPlayer:
		var cmd tea.Cmd
		m.playerView, cmd = m.playerView.Update(msg)
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
			return m, m.listView.Init()
		case viewPlayer:
			return m, m.playerView.Init()
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.viewState {
	case viewConnecting:
		return m.connectingView.View()
	case viewList:
		return m.listView.View()
	case viewPlayer:
		return m.playerView.View()
	}
	return ""
}

func InitialModel() model {
	return model{
		viewState:      viewConnecting,
		connectingView: NewConnectingView(),
		listView:       NewListView(),
		playerView:     NewPlayerView(),
		inits: map[viewState]bool{
			viewConnecting: false,
			viewList:       false,
			viewPlayer:     false,
		},
	}
}
