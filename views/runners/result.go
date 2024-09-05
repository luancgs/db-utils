package runners

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luancgs/db-utils/views/components"
)

func ResultRunner(text string, status int) {
	p := tea.NewProgram(components.ResultModelInit(text, status))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
