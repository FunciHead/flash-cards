package cli

import (
	"flash-cards/flashcards"
)

const pathToData = "data/deck.csv"

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

func getTheAnswers(f flashcards.FlashCards) []string {
	answers := []string{}
	for _, item := range f.MyFlashCards {
		answers = append(answers, item.Answer)
	}
	return answers
}

func setQuestionUpdate(flashs flashcards.FlashCards, question string, index int) {
	flashs.MyFlashCards[index].Question = question
	flashcards.SaveFlashcards(pathToData, flashs.MyFlashCards)
}

func setAnswerUpdate(flashs flashcards.FlashCards, answer string, index int) {
	flashs.MyFlashCards[index].Answer = answer
	flashcards.SaveFlashcards(pathToData, flashs.MyFlashCards)
}

func flahsCardDelete(flashs flashcards.FlashCards, index int) {
	flashs.MyFlashCards = append(flashs.MyFlashCards[:index], flashs.MyFlashCards[index+1:]...)
	flashcards.SaveFlashcards(pathToData, flashs.MyFlashCards)
}
