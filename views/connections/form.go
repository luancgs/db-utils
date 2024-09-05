package connections

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	username = iota
	password
	host
	port
	database
	query
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type ConnectionFormModel struct {
	PreviousChoices []string
	Inputs          []textinput.Model
	Focused         int
	Err             error
}

func expValidator(s string) error {
	// The 3 character should be a slash (/)
	// The rest should be numbers
	e := strings.ReplaceAll(s, "/", "")
	_, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fmt.Errorf("EXP is invalid")
	}

	// There should be only one slash and it should be in the 2nd index (3rd character)
	if len(s) >= 3 && (strings.Index(s, "/") != 2 || strings.LastIndex(s, "/") != 2) {
		return fmt.Errorf("EXP is invalid")
	}

	return nil
}

func portValidator(s string) error {
	_, err := strconv.ParseInt(s, 10, 64)
	return err
}

func ConnectionFormModelInit(previousChoices []string) ConnectionFormModel {
	var inputs []textinput.Model = make([]textinput.Model, 6)

	inputs[host] = textinput.New()
	inputs[host].Placeholder = "localhost"
	inputs[host].Focus()
	inputs[host].CharLimit = 253
	inputs[host].Width = 70
	inputs[host].Prompt = ""

	inputs[port] = textinput.New()
	inputs[port].Placeholder = "5432"
	inputs[port].CharLimit = 5
	inputs[port].Width = 5
	inputs[port].Prompt = ""
	inputs[port].Validate = portValidator

	inputs[username] = textinput.New()
	inputs[username].Placeholder = "username"
	inputs[username].CharLimit = 40
	inputs[username].Width = 40
	inputs[username].Prompt = ""

	inputs[password] = textinput.New()
	inputs[password].Placeholder = "P4$$w0Rd"
	inputs[password].CharLimit = 40
	inputs[password].Width = 30
	inputs[password].Prompt = ""

	inputs[database] = textinput.New()
	inputs[database].Placeholder = "database_name"
	inputs[database].CharLimit = 20
	inputs[database].Width = 20
	inputs[database].Prompt = ""

	return ConnectionFormModel{
		PreviousChoices: previousChoices,
		Inputs:          inputs,
		Focused:         0,
	}
}

func (m ConnectionFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ConnectionFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.Inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.Focused == len(m.Inputs)-1 {
				return m, tea.Quit
			}

			m.nextInput()
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "shift+tab", "ctrl+p", "up":
			m.prevInput()
		case "tab", "ctrl+n", "down":
			m.nextInput()
		}

		for i := range m.Inputs {
			m.Inputs[i].Blur()
		}

		m.Inputs[m.Focused].Focus()

	// We handle errors just like any other message
	case error:
		m.Err = msg
		return m, nil
	}

	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m ConnectionFormModel) View() string {
	return fmt.Sprintf(
		` Total: $21.50:

 %s
 %s

 %s          %s
 %s  %s

 %s          %s
 %s  %s
 %s
`,
		inputStyle.Width(30).Render("Host"),
		m.Inputs[host].View(),
		inputStyle.Width(8).Render("Database"),
		inputStyle.Width(6).Render("Port"),
		m.Inputs[database].View(),
		m.Inputs[port].View(),
		inputStyle.Width(8).Render("Username"),
		inputStyle.Width(8).Render("Password"),
		m.Inputs[username].View(),
		m.Inputs[password].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

func (m *ConnectionFormModel) nextInput() {
	m.Focused = (m.Focused + 1) % len(m.Inputs)
}

func (m *ConnectionFormModel) prevInput() {
	m.Focused--

	if m.Focused < 0 {
		m.Focused = len(m.Inputs) - 1
	}
}
