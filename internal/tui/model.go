package tui

import (
	e "github.com/alexlangev/mfp/internal/episodes"
)

type Model struct {
	isConnected     bool
	episodes        e.Episodes
	selectedEpisode int
}

func InitialModel() Model {
	return Model{
		isConnected: false,
	}
}
