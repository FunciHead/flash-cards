package cli

func updateMenu(m *Model) {
	if m.inMenu {
		if m.cursor == 0 {
			m.mode = UPDATEQUESTIONS
		} else if m.cursor == 1 {
			m.mode = UPDATEANSWERS
		}
		m.inMenu = false
	} else {
		m.inMenu = true
	}
}

func changeAQ(m *Model) {
	if m.inSubMenu {
		if m.mode == UPDATEANSWERS {
			changeAnswer(m)
		} else {
			changeQuestion(m)
		}
		m.inSubMenu = false
	} else {
		m.inSubMenu = true
	}
}

func changeAnswer(m *Model) {
	answer := m.choices["updateanswers"][m.cursor]
	answer += "!"
	setAnswerUpdate(initialize(), answer, m.cursor)
	updateValuesAnswers(m)
	m.mode = UPDATE
	m.cursor = 1
	m.inMenu = true
}

func changeQuestion(m *Model) {
	question := m.choices["updatequestions"][m.cursor]
	question += "!" //temporary solution
	setQuestionUpdate(initialize(), question, m.cursor)
	updateValuesQuestions(m)
	m.mode = UPDATE
	m.cursor = 0
	m.inMenu = true
}

func updateValuesQuestions(m *Model) {
	flash := initialize()
	m.choices["delete"] = getTheQuestions(flash)
	m.choices["updatequestions"] = getTheQuestions(flash)
}

func updateValuesAnswers(m *Model) {
	flash := initialize()
	m.choices["updateanswers"] = getTheAnswers(flash)
}
