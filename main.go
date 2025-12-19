package main

import (
	"fmt"
	"os"

	"github.com/MrShanks/kangoo/kanban"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initialModel := kanban.New()

	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
