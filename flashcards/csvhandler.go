package flashcards

import (
	"encoding/csv"
	"os"
)

type FlashCard struct {
	Question     string
	Answer       string
	RightOrWrong string
}

func LoadFlashcards(path string) ([]FlashCard, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var cards []FlashCard
	for _, record := range records[1:] {
		if len(record) >= 2 {
			cards = append(cards, FlashCard{
				Question:     record[0],
				Answer:       record[1],
				RightOrWrong: record[2],
			})
		}
	}
	return cards, nil
}

func SaveFlashcards(path string, cards []FlashCard) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Question", "Answer", "RightOrWrong"})

	for _, card := range cards {
		writer.Write([]string{card.Question, card.Answer, card.RightOrWrong})
	}
	return nil
}

func EnsureCSVFile(path string) error {
	os.MkdirAll("data", os.ModeAppend)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}

		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		writer.Write([]string{"Question", "Answer", "RightOrWrong"})

	}
	return nil
}
