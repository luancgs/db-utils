package general

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luancgs/db-utils/config"
	"github.com/luancgs/db-utils/views/utils"
)

type ConnectionSelectionModel struct {
	PreviousChoices []string
	Choices         []string
	Cursor          int
	Selected        string
	Err             error
}

func ConnectionSelectionModelInit(previousChoices []string) ConnectionSelectionModel {
	return ConnectionSelectionModel{
		PreviousChoices: previousChoices,
		Choices:         config.GetConnectionsEnum().ToArray(),
		Selected:        "",
	}
}

func (m ConnectionSelectionModel) Init() tea.Cmd {
	return nil
}

func (m ConnectionSelectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		case "enter", " ":
			m.Selected = m.Choices[m.Cursor]
			return m, tea.Quit
		}

	case error:
		m.Err = msg
	}

	return m, nil
}

func (m ConnectionSelectionModel) View() string {
	output := utils.PreviousChoicesString(m.PreviousChoices)
	output += "Choose a connection format:\n\n"

	for i, choice := range m.Choices {

		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		if choice == m.Selected {
			cursor = "✔"
		}

		output += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	output += "\nPress q to quit.\n"

	return output
}
