package main

import (
	"flash-cards/flashcards"
	"fmt"
)

const pathToData = "data/deck.csv"

func main() {
	flashcards.EnsureCSVFile(pathToData)
	fmt.Println("Oi mundo!")
}
