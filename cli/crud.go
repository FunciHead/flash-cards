package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

const MENU = -1
const READ = 0
const CREATE = 1
const UPDATE = 2
const DELETE = 3
const UPDATEQUESTIONS = 4
const UPDATEANSWERS = 5

type Model struct {
	choices   map[string][]string
	cursor    int
	mode      int
	inMenu    bool
	inSubMenu bool
}

func InitialModel() Model {
	return Model{
		cursor: 0,
		mode:   -1,
		choices: map[string][]string{
			"menu":            {"Start", "Create new flash card", "Edit flash cards", "Delete flash cards"},
			"read":            {"Start"},
			"update":          {"Update question", "Update answer"},
			"delete":          getTheQuestions(initialize()),
			"create":          {"Create new flash card"},
			"updatequestions": getTheQuestions(initialize()),
			"updateanswers":   getTheAnswers(initialize()),
		},
		inMenu:    false,
		inSubMenu: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			if m.mode == -1 {
				return m, tea.Quit
			} else {
				returnButton(&m)
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < m.lenOfMenu() {
				m.cursor++
			}
		case "enter":
			if m.mode == MENU {
				mainMenu(&m)
			}
			if m.mode == UPDATE {
				updateMenu(&m)
			}
			if m.mode == UPDATEANSWERS || m.mode == UPDATEQUESTIONS {
				changeAQ(&m)
			}
			if m.mode == DELETE {
				deleteMenu(&m)
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	switch m.mode {
	case MENU:
		return menuUpdateView(&m, "menu")
	case READ:
		return menuUpdateView(&m, "read")
	case UPDATE:
		return menuUpdateView(&m, "update")
	case CREATE:
		return menuUpdateView(&m, "create")
	case DELETE:
		return menuUpdateView(&m, "delete")
	case UPDATEQUESTIONS:
		return menuUpdateView(&m, "updatequestions")
	case UPDATEANSWERS:
		return menuUpdateView(&m, "updateanswers")
	}
	return menuUpdateView(&m, "menu")
}
