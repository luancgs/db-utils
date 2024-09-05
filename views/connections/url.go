package connections

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/luancgs/db-utils/views/utils"
)

type ConnectionURLModel struct {
	previousChoices []string
	textInput       textinput.Model
	err             error
}

func ConnectionURLModelInit(previousChoices []string) ConnectionURLModel {
	ti := textinput.New()
	ti.Placeholder = "protocol://user:password@host:port/db_name?schema=public"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 70

	return ConnectionURLModel{
		previousChoices: previousChoices,
		textInput:       ti,
		err:             nil,
	}
}

func (m ConnectionURLModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ConnectionURLModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m ConnectionURLModel) View() string {
	output := utils.PreviousChoicesString(m.previousChoices)
	output += fmt.Sprintf("Connection URL:\n%s\n", m.textInput.View())

	return output
}
