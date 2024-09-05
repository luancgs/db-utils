package utils

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	CHOICES_ICON = "󰣏"
)

func PreviousChoicesString(previousChoices []string) string {
	return fmt.Sprintf("%s %s\n\n", CHOICES_ICON, strings.Join(previousChoices, "\n"))
}

func ProgramRunner(program tea.Model) {
	p := tea.NewProgram(program)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
