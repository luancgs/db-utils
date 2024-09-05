package runners

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luancgs/db-utils/views/components"
)

func LoadingRunner(startText string, textChan chan string, doneChan chan bool) {
	p := tea.NewProgram(components.LoadingModelInitWithChannels(startText, textChan, doneChan))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
