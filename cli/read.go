package cli

import (
	"flash-cards/flashcards"
	"math/rand/v2"
)

func playFlashCards(m *Model) {

	if m.inMenu {
		m.mode = GAME
		m.index = 0
		m.inMenu = false
		m.flashcardsArray = getInformation(m)

	} else {
		m.inMenu = true
	}
}

func play(m *Model) {
	output(m)
	if m.inSubMenu {
		switch m.cursor {
		case 0:
			show(m)
		case 1:
			hide(m)
		case 2:
			next(m)
			output(m)

		case 3:
			previous(m)
			output(m)

		}
	} else {
		m.inSubMenu = true
	}

}
func getInformation(m *Model) []flashcards.FlashCard {
	flash := initializeAndRandomizeFlashcards()
	m.secretString = flash[m.index].Question
	return flash
}

func output(m *Model) {
	m.secretString = m.flashcardsArray[m.index].Question
}

func show(m *Model) {
	m.secretString += ":  "
	m.secretString += m.flashcardsArray[m.index].Answer
}

func hide(m *Model) {
	m.secretString = m.flashcardsArray[m.index].Question
}

func next(m *Model) {
	if m.index < len(m.flashcardsArray)-1 {
		m.index++
	} else {
		m.mode = READ
		m.cursor = 0
		m.inSubMenu = false
		m.inMenu = true
		m.index = 0
	}
}

func previous(m *Model) {
	if m.index > 0 {
		m.index--
	}
}

func initializeAndRandomizeFlashcards() []flashcards.FlashCard {
	flash := initialize().MyFlashCards
	for i := len(flash) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		flash[i], flash[j] = flash[j], flash[i]
	}
	return flash
}
