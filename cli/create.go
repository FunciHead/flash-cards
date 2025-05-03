package cli

func CreateFlashCard(m *Model) {
	if m.inMenu {
		create(m)
		m.inMenu = false
	} else {
		m.inMenu = true
	}
}

func create(m *Model) {
	flash := initialize()
	question := "la"
	answer := "polizia"
	addFlashCard(flash, question, answer)
	flash = initialize()
	m.choices["delete"] = getTheQuestions(flash)
	m.choices["updatequestions"] = getTheQuestions(flash)
	m.choices["updateanswers"] = getTheQuestions(flash)
	m.cursor = CREATE
	m.mode = MENU

}
