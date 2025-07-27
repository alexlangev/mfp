package tui

import (
	e "github.com/alexlangev/mfp/internal/episodes"
)

type Model struct {
	isConnected     bool
	episodes        e.Episodes
	selectedEpisode int
	err             error
}

type errMsg struct{ err error }

// implement the error interface on the message
func (e errMsg) Error() string { return e.err.Error() }

type epsMsg struct{ episodes e.Episodes }

func InitialModel() Model {
	return Model{
		isConnected: false,
	}
}
