package main

import (
	"flash-cards/cli"
	"flash-cards/flashcards"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const pathToData = "data/deck.csv"

func main() {
	flashcards.EnsureCSVFile(pathToData)
	p := tea.NewProgram(cli.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
