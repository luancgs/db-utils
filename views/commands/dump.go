package commands

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

type DumpCommandModel struct {
	spinner spinner.Model
	Text    string
	Err     error
}

func DumpCommandModelInit() DumpCommandModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return DumpCommandModel{
		spinner: s,
		Text:    "Dumping database...",
		Err:     nil,
	}
}

func (m DumpCommandModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m DumpCommandModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			os.Exit(0)
			return m, tea.Quit
		default:
			return m, nil
		}

	case error:
		m.Err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m DumpCommandModel) View() string {
	if m.Err != nil {
		return m.Err.Error()
	}

	output := fmt.Sprintf("\n%s%s\n\n", m.spinner.View(), textStyle(m.Text))

	return output
}
