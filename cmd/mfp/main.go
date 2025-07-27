package main

import (
	"fmt"
	"os"

	"github.com/alexlangev/mfp/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// fmt.Println("Hello there!")
	// fmt.Println()
	// x, _ := episodes.GetEpisodes()
	// fmt.Println(x[0])
	// fmt.Println(x[2])
	// fmt.Println(x[69])
	// fmt.Println(x[74])

	if _, err := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
