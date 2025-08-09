package main

import (
	"fmt"
	"os"

	"github.com/alexlangev/mfp/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	_, err := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
