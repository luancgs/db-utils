package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/luancgs/db-utils/databases"
	"github.com/luancgs/db-utils/errors"
	"github.com/luancgs/db-utils/views/runners"
)

func NewRestoreCommand() *RestoreCommand {
	restoreCommand := &RestoreCommand{
		flagSet: flag.NewFlagSet("restore", flag.ContinueOnError),
	}

	restoreCommand.flagSet.BoolVar(&restoreCommand.help, "help", false, "Show help message")
	restoreCommand.flagSet.BoolVar(&restoreCommand.help, "h", false, "Show help message")

	restoreCommand.flagSet.StringVar(&restoreCommand.databaseUrl, "url", "", "Database URL")
	restoreCommand.flagSet.StringVar(&restoreCommand.databaseUrl, "u", "", "Database URL")

	restoreCommand.flagSet.StringVar(&restoreCommand.restoreFile, "file", "", "File to restore the database")
	restoreCommand.flagSet.StringVar(&restoreCommand.restoreFile, "f", "", "File to restore the database")

	return restoreCommand
}

type RestoreCommand struct {
	flagSet     *flag.FlagSet
	help        bool
	databaseUrl string
	restoreFile string
}

func (pc *RestoreCommand) Name() string {
	return pc.flagSet.Name()
}

func (pc *RestoreCommand) Init(args []string) error {
	return pc.flagSet.Parse(args)
}

func (pc *RestoreCommand) Run() {
	if pc.help {
		fmt.Println(pc.Help())
		os.Exit(0)
	}

	if pc.databaseUrl == "" || pc.restoreFile == "" {
		runners.ResultRunner("Database URL and restore file are required", 0)
		return
	}

	db, err := databases.ParseUrl(pc.databaseUrl)
	errors.ErrorHandler("Error while parsing database URL", err)

	ok, err := db.Restore(pc.restoreFile)
	errors.ErrorHandler("Error while restoring from file", err)

	if !ok {
		runners.ResultRunner("Database restore failed!", 0)
		return
	}

	runners.ResultRunner("Database restored successfully.", 1)
}

func (pc RestoreCommand) Help() string {
	var output string

	output += fmt.Sprintln("Usage: db-utils dump [flags]")
	output += fmt.Sprintln()
	output += fmt.Sprintln("Flags:")
	output += fmt.Sprintln("  -h, --help     Show help message")
	output += fmt.Sprintln("  -u, --url      Database URL")
	output += fmt.Sprintln("  -f, --file     File to restore the database")

	return output
}
