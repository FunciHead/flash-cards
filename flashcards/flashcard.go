package flashcards

//tea "github.com/charmbracelet/bubbletea"

type Model struct {
	FlashCardFace bool
	RightAnswers  uint
	MyFlashCards  []FlashCard
}

/*
func (m Model) Init() tea.Cmd {
	return nil
}
*/
