package flashcards

const pathToData = "data/deck.csv"

type FlashCards struct {
	FlashCardFace bool
	RightAnswers  uint
	MyFlashCards  []FlashCard
}

func InitializeFlashcards() FlashCards {
	flashs, _ := LoadFlashcards(pathToData)

	return FlashCards{
		FlashCardFace: false,
		RightAnswers:  0,
		MyFlashCards:  flashs,
	}
}
