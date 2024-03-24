package commands

import (
	"flag"
	"fmt"

	"github.com/luancgs/db-utils/databases"
)

func NewRestoreCommand() *RestoreCommand {
	restoreCommand := &RestoreCommand{
		flagSet: flag.NewFlagSet("restore", flag.ContinueOnError),
	}

	restoreCommand.flagSet.StringVar(&restoreCommand.databaseUrl, "url", "", "url of the database")
	restoreCommand.flagSet.StringVar(&restoreCommand.restoreFile, "file", "", "file to restore the database")

	return restoreCommand
}

type RestoreCommand struct {
	flagSet     *flag.FlagSet
	databaseUrl string
	restoreFile string
}

func (rc *RestoreCommand) Name() string {
	return rc.flagSet.Name()
}

func (rc *RestoreCommand) Init(args []string) error {
	return rc.flagSet.Parse(args)
}

func (rc *RestoreCommand) Run() error {
	if rc.databaseUrl == "" {
		return fmt.Errorf("database url is required")
	}

	db, err := databases.ParseUrl(rc.databaseUrl)
	if err != nil {
		return fmt.Errorf("failed parsing database url: %w", err)
	}

	ok, err := db.Restore(rc.restoreFile)
	if err != nil {
		return fmt.Errorf("failed restoring database: %w", err)
	}

	if !ok {
		return fmt.Errorf("failed restoring database")
	}

	fmt.Println("Database restored successfully.")

	return nil
}
