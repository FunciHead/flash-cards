package cli

func deleteMenu(m *Model) {
	if m.inMenu {
		deleteEntry(m)
		m.inMenu = false
	} else {
		m.inMenu = true
	}

}

func deleteEntry(m *Model) {
	option := m.cursor
	flahsCardDelete(initialize(), option)
	flash := initialize()

	m.choices["delete"] = getTheQuestions(flash)
	m.choices["updatequestions"] = getTheQuestions(flash)
	m.choices["updateanswers"] = getTheQuestions(flash)
	m.cursor = DELETE
	m.mode = MENU
}
