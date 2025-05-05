package cli

import "fmt"

func menuUpdateView(m *Model, key string) string {
	s := "Flash card CLI\n"
	if !m.inputing {
		if m.mode == GAME {
			s += "   "
			s += m.secretString
			s += "\n"
		}
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
	} else {
		m.textInput.Focus()
		s += fmt.Sprintf("Type something and press enter: \n\n%s\n\n", m.textInput.View())
		if m.submitted != "" {
			s += fmt.Sprintf("you entered: %s", m.submitted)
		}
	}
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
	case UPDATEQUESTIONS:
		length = len(m.choices["updatequestions"])
	case UPDATEANSWERS:
		length = len(m.choices["updateanswers"])
	case GAME:
		length = len(m.choices["game"])
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
	if m.mode == UPDATEANSWERS || m.mode == UPDATEQUESTIONS {
		switch m.mode {
		case UPDATEQUESTIONS:
			m.cursor = 0
		case UPDATEANSWERS:
			m.cursor = 1
		default:
			m.cursor = 0
		}
		m.mode = UPDATE
		m.inMenu = true
		m.inSubMenu = false
	} else if m.mode == GAME {
		m.cursor = 0
		m.mode = READ
		m.inMenu = true
		m.inSubMenu = false
		m.index = 0
	} else {
		m.cursor = m.mode
		m.mode = MENU
		m.inMenu = false
	}
}
