package main

import (
	"fmt"
	"os"

	"github.com/luancgs/db-utils/commands"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must provide a subcommand")
	}

	cmds := []Runner{
		commands.NewCloneCommand(),
		commands.NewDumpCommand(),
		commands.NewRestoreCommand(),
		commands.NewQueryCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}

func main() {
	if _, err := os.Stat("/var/run/docker.sock"); os.IsNotExist(err) {
		fmt.Println("Docker is not running. Please start docker and try again.")
		os.Exit(1)
	}

	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
