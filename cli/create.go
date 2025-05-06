package cli

func CreateFlashCard(m *Model) {
	if m.inMenu {
		m.inputing = true
		m.textInput.Focus()
		m.textInput.SetValue("")
		m.submitted = "" // Reset submitted value
	} else {
		m.inMenu = true
	}
}

func create(m *Model) {
	flash := initialize()

	// If we haven't collected the question yet
	if m.submitted == "" {
		m.submitted = m.textInput.Value()
		m.textInput.SetValue("")
		return
	}

	// If we have the question, now collect the answer
	question := m.submitted
	answer := m.textInput.Value()
	m.textInput.SetValue("")

	addFlashCard(flash, question, answer)
	flash = initialize()
	m.choices["delete"] = getTheQuestions(flash)
	m.choices["updatequestions"] = getTheQuestions(flash)
	m.choices["updateanswers"] = getTheAnswers(flash)
	m.cursor = CREATE
	m.mode = MENU
	m.inMenu = false
	m.inputing = false
	m.submitted = "" // Reset for next time
}
