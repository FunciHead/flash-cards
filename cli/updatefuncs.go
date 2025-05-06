package cli

// updateMenu handles the initial selection between updating questions or answers
func updateMenu(m *Model) {
	if m.inMenu {
		// Set the mode based on user selection (0 for questions, 1 for answers)
		if m.cursor == 0 {
			m.mode = UPDATEQUESTIONS
		} else {
			m.mode = UPDATEANSWERS
		}
		m.inMenu = false
		m.inSubMenu = true
		m.cursor = 0 // Reset cursor to start of list
	} else {
		m.inMenu = true
	}
}

// changeAQ handles the selection of which item to edit
func changeAQ(m *Model) {
	if !m.inSubMenu {
		m.inSubMenu = true
		return
	}

	if !m.inputing {
		m.inputing = true
		m.textInput.Focus()

		if m.mode == UPDATEQUESTIONS {
			m.textInput.SetValue(m.choices["updatequestions"][m.cursor])
		} else {
			m.textInput.SetValue(m.choices["updateanswers"][m.cursor])
		}
		return
	}

	// Save changes
	flash := initialize()
	if m.mode == UPDATEQUESTIONS {
		setQuestionUpdate(flash, m.textInput.Value(), m.cursor)
		updateValuesQuestions(m)
		m.mode = UPDATE
		m.cursor = 0
	} else {
		setAnswerUpdate(flash, m.textInput.Value(), m.cursor)
		updateValuesAnswers(m)
		m.mode = UPDATE
		m.cursor = 1
	}

	// Reset states
	m.inMenu = true
	m.inSubMenu = false
	m.inputing = false
	m.textInput.SetValue("")
}

// updateValuesQuestions updates the question-related choices in the menu
func updateValuesQuestions(m *Model) {
	flash := initialize()
	m.choices["delete"] = getTheQuestions(flash)
	m.choices["updatequestions"] = getTheQuestions(flash)
}

// updateValuesAnswers updates the answer-related choices in the menu
func updateValuesAnswers(m *Model) {
	flash := initialize()
	m.choices["updateanswers"] = getTheAnswers(flash)
}
