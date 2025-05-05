package cli

func CreateFlashCard(m *Model) {
	if m.inMenu {
		m.inputing = true
		create(m)
	} else {
		m.inMenu = true
	}
}

func create(m *Model) {
	flash := initialize()
	m.submitted = m.textInput.Value()
	question := m.submitted
	m.textInput.SetValue("")
	answer := "tu"
	addFlashCard(flash, question, answer)
	flash = initialize()
	m.choices["delete"] = getTheQuestions(flash)
	m.choices["updatequestions"] = getTheQuestions(flash)
	m.choices["updateanswers"] = getTheAnswers(flash)
	m.cursor = CREATE
	m.mode = MENU
	m.inMenu = false

}
