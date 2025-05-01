package cli

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

const MENU = -1
const READ = 0
const CREATE = 1
const UPDATE = 2
const DELETE = 3

type Model struct {
	choices map[string][]string
	cursor  int
	mode    int
}

func InitialModel() Model {
	return Model{
		cursor: 0,
		mode:   -1,
		choices: map[string][]string{
			"menu":   {"Start", "Create new flash card", "Edit flash cards", "Delete flash cards"},
			"read":   {"Start"},
			"update": getTheQuestions(initialize()),
			"delete": getTheQuestions(initialize()),
			"create": {"Create new flash card"},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.mode == -1 {
				return m, tea.Quit
			} else {
				m.mode = -1
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-2 {
				m.cursor++
			}
		case "enter":
			m.changeMode(&m)
		}
	}
	return m, nil
}

func (m Model) View() string {
	switch m.mode {
	case MENU:
		return menu(&m, "menu")
	case READ:
		return menu(&m, "read")
	case UPDATE:
		return menu(&m, "update")
	case CREATE:
		return menu(&m, "create")
	case DELETE:
		return menu(&m, "delete")
	}
	return "p"
}

func menu(m *Model, key string) string {
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
