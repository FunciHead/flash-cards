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
	case UPDATEQUESTIONS:
		length = len(m.choices["updatequestions"])
	case UPDATEANSWERS:
		length = len(m.choices["updateanswers"])

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
	if m.mode > 3 {
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
	} else {
		m.cursor = m.mode
		m.mode = MENU
		m.inMenu = false
	}
}
