package commands

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/luancgs/db-utils/views/utils"
)

type RunQueryModel struct {
	previousChoices []string
	textInput       textinput.Model
	err             error
}

func RunQueryModelInit(previousChoices []string) RunQueryModel {
	ti := textinput.New()
	ti.Placeholder = "protocol://user:password@host:port/db_name?schema=public"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 70

	return RunQueryModel{
		previousChoices: previousChoices,
		textInput:       ti,
		err:             nil,
	}
}

func (m RunQueryModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m RunQueryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m RunQueryModel) View() string {
	return fmt.Sprintf(
		"%s What’s your favorite Pokémon?\n\n%s\n\n%s",
		utils.PreviousChoicesString(m.previousChoices),
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
