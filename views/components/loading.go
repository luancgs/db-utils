package components

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LoadingModel struct {
	doneChan   chan bool
	textChan   chan string
	Text       string
	textRender func(strs ...string) string
	spinner    spinner.Model
	Err        error
}

type loadingModelInitParams struct {
	spinnerModel spinner.Spinner
	spinnerColor string
	textColor    string
	startText    string
	textChan     chan string
	doneChan     chan bool
}

func loadingModelSetup(params loadingModelInitParams) LoadingModel {
	s := spinner.New()
	s.Spinner = params.spinnerModel
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(params.spinnerColor))

	return LoadingModel{
		spinner:    s,
		textChan:   params.textChan,
		doneChan:   params.doneChan,
		textRender: lipgloss.NewStyle().Foreground(lipgloss.Color(params.textColor)).Render,
		Text:       params.startText,
		Err:        nil,
	}
}

func LoadingModelInit(text string) LoadingModel {
	return loadingModelSetup(loadingModelInitParams{
		spinnerModel: spinner.Dot,
		spinnerColor: "69",
		textColor:    "252",
		startText:    text,
	})
}

func LoadingModelInitWithChannels(text string, textChan chan string, doneChan chan bool) LoadingModel {
	return loadingModelSetup(loadingModelInitParams{
		spinnerModel: spinner.Dot,
		spinnerColor: "69",
		textColor:    "252",
		startText:    text,
		textChan:     textChan,
		doneChan:     doneChan,
	})
}

func LoadingModelInitCustom(params loadingModelInitParams) LoadingModel {
	return loadingModelSetup(params)
}

func (m LoadingModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m LoadingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	select {
	case <-m.doneChan:
		return m, tea.Quit
	case text, ok := <-m.textChan:
		if ok {
			m.Text = text
		}
	default:

	}

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

func (m LoadingModel) View() string {
	if m.Err != nil {
		return m.Err.Error()
	}

	output := fmt.Sprintf("\n%s%s\n\n", m.spinner.View(), m.textRender(m.Text))

	return output
}
