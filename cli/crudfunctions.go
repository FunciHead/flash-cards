package cli

import "fmt"

func menuUpdateView(m *Model, key string) string {
	s := "Flash card CLI\n"
	if values, exists := m.choices[key]; exists {
		for i, choices := range values {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s  %s\n", cursor, choices)
		}
	}
	s += "Press \"q\" to quit"
	return s
}

func (m Model) changeMode(me *Model) {
	me.mode = m.cursor
}

func (m Model) lenOfMenu() int {
	var length int = 0
	switch m.mode {
	case MENU:
		length = len(m.choices["menu"])
	case READ:
		length = len(m.choices["read"])
	case UPDATE:
		length = len(m.choices["update"])
	case CREATE:
		length = len(m.choices["create"])
	case DELETE:
		length = len(m.choices["delete"])

	}
	return length - 1
}

func mainMenu(m *Model) {
	m.changeMode(m)
	if m.mode != MENU {
		m.cursor = 0
	}
}

func returnButton(m *Model) {

	m.cursor = m.mode
	m.mode = -1
	m.inMenu = false
}

func updateMenu(m *Model) {
	if m.inMenu {
		if m.cursor == 0 {
			m.mode = UPDATEQUESTIONS
		} else {
			m.mode = UPDATEANSWERS
		}
		m.inMenu = false
	} else {
		m.inMenu = true
	}
}

func changeAnswer(m *Model) {
	answer := m.choices["delete"][m.cursor]
	answer += "zika"
	setQuestion(initialize(), answer, m.cursor)
	updateValuesAnswers(m)
	m.mode = -1
	m.inMenu = false
}

func changeQuestion(m *Model) {
	question := m.choices["delete"][m.cursor]
	question += "!" //temporary solution
	setQuestion(initialize(), question, m.cursor)
	updateValuesQuestions(m)
	m.mode = -1
	m.inMenu = false
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
