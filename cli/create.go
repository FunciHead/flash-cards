package cli

func CreateFlashCard(m *Model) {
	if m.inMenu {
		create(m)
	} else {
		m.inMenu = true
	}
}

func create(m *Model) {
	flash := initialize()
	//ti := textinput.New()
	question := "la"
	answer := "polizia"
	addFlashCard(flash, question, answer)
	flash = initialize()
	m.choices["delete"] = getTheQuestions(flash)
	m.choices["updatequestions"] = getTheQuestions(flash)
	m.choices["updateanswers"] = getTheAnswers(flash)
	m.cursor = CREATE
	m.mode = MENU
	m.inMenu = false

}
