package views

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luancgs/db-utils/config"
	"github.com/luancgs/db-utils/views/connections"
	"github.com/luancgs/db-utils/views/general"
)

func InterativeFlow() {
	choices := []string{}

	p := tea.NewProgram(general.DBMSSelectionModelInit(choices))

	model, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	dbmsSelection, ok := model.(general.DBMSSelectionModel)

	if !ok || dbmsSelection.Err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	choices = append(choices, dbmsSelection.Selected)

	p = tea.NewProgram(general.ConnectionSelectionModelInit(choices))

	model, err = p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	connectionSelection, ok := model.(general.ConnectionSelectionModel)

	if !ok || connectionSelection.Err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	choices = append(choices, connectionSelection.Selected)

	if connectionSelection.Selected == config.GetConnectionsEnum().ConnectionURL {
		runConnectionURL(&choices)
	} else {
		runConnectionForm(&choices)
	}

	p = tea.NewProgram(general.CommandSelectionModelInit(choices))

	model, err = p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	commandSelection, ok := model.(general.CommandSelectionModel)

	if !ok || commandSelection.Err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func runConnectionURL(choices *[]string) {
	p := tea.NewProgram(connections.ConnectionURLModelInit(*choices))

	model, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	_, ok := model.(connections.ConnectionURLModel)

	if !ok {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	*choices = append(*choices, "Connected: true")
}

func runConnectionForm(choices *[]string) {
	p := tea.NewProgram(connections.ConnectionFormModelInit(*choices))

	model, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	connectionForm, ok := model.(connections.ConnectionFormModel)

	if !ok || connectionForm.Err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	*choices = append(*choices, "Connected: true")
}
