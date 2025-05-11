package cli

// updateMenu handles the initial selection between updating questions or answers
func updateMenu(m *Model) {
	if m.inMenu {
		if m.cursor == 0 {
			m.mode = UPDATEQUESTIONS
			m.choices["updatequestions"] = getTheQuestions(initialize())
		} else if m.cursor == 1 {
			m.mode = UPDATEANSWERS
			m.choices["updateanswers"] = getTheAnswers(initialize())
		}
		m.inMenu = false
	} else {
		m.inMenu = true
	}
}

// changeAQ handles the selection of which item to edit
func changeAQ(m *Model) {
	if m.inSubMenu {
		if m.mode == UPDATEANSWERS {
			changeAnswer(m)
		} else {
			changeQuestion(m)
		}
	} else {
		m.inSubMenu = true
	}
}

func changeAnswer(m *Model) {
	if !m.inputing {
		m.inputing = true
		m.textInput.Focus()
		answer := m.choices["updateanswers"][m.cursor]
		m.textInput.SetValue(answer)
		return
	}

	setAnswerUpdate(initialize(), m.textInput.Value(), m.cursor)
	updateValuesAnswers(m)
	m.mode = UPDATE
	m.cursor = 1
	m.inMenu = true
	m.inputing = false
	m.textInput.SetValue("")
	m.inSubMenu = false

}

func changeQuestion(m *Model) {
	if !m.inputing {
		m.inputing = true
		m.textInput.Focus()
		question := m.choices["updatequestions"][m.cursor]
		m.textInput.SetValue(question)
		return
	}

	setQuestionUpdate(initialize(), m.textInput.Value(), m.cursor)
	updateValuesQuestions(m)
	m.mode = UPDATE
	m.cursor = 0
	m.inMenu = true
	m.inputing = false
	m.textInput.SetValue("")
	m.inSubMenu = false

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
