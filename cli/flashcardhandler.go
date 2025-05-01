package cli

import (
	"flash-cards/flashcards"
)

func initialize() flashcards.FlashCards {
	return flashcards.InitializeFlashcards()
}

func getTheQuestions(f flashcards.FlashCards) []string {
	questions := []string{}
	for _, item := range f.MyFlashCards {
		questions = append(questions, item.Question)
	}
	return questions
}
