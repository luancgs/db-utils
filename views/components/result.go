package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	err = iota
	success
)

type ResultModel struct {
	status     int
	textRender func(strs ...string) string
	Text       string
}

type resultModelInitParams struct {
	status int
	text   string
}

func resultModelSetup(params resultModelInitParams) ResultModel {

	var textColor string

	switch params.status {
	case success:
		textColor = "86"
	default:
		textColor = "161"
	}

	return ResultModel{
		Text:       params.text,
		textRender: lipgloss.NewStyle().Foreground(lipgloss.Color(textColor)).Render,
		status:     params.status,
	}
}

func ResultModelInit(text string, status int) ResultModel {
	return resultModelSetup(resultModelInitParams{
		status: status,
		text:   text,
	})
}

func (m ResultModel) Init() tea.Cmd {
	return nil
}

func (m ResultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.View()
	return m, tea.Quit
}

func (m ResultModel) View() string {
	var lineStart string

	switch m.status {
	case success:
		lineStart = "✓ "
	case err:
		lineStart = "✗ "
	default:
		lineStart = ""
	}

	output := fmt.Sprintf("\n%s%s\n\n", lineStart, m.textRender(m.Text))

	return output
}
