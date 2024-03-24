package commands

import (
	"flag"
	"fmt"

	"github.com/luancgs/db-utils/databases"
)

func NewPopulateCommand() *PopulateCommand {
	dumpCommand := &PopulateCommand{
		flagSet: flag.NewFlagSet("populate", flag.ContinueOnError),
	}

	dumpCommand.flagSet.StringVar(&dumpCommand.databaseUrl, "url", "", "url of the database")
	dumpCommand.flagSet.StringVar(&dumpCommand.populateFile, "file", "", "file to populate the database")

	return dumpCommand
}

type PopulateCommand struct {
	flagSet      *flag.FlagSet
	databaseUrl  string
	populateFile string
}

func (pc *PopulateCommand) Name() string {
	return pc.flagSet.Name()
}

func (pc *PopulateCommand) Init(args []string) error {
	return pc.flagSet.Parse(args)
}

func (pc *PopulateCommand) Run() error {
	if pc.databaseUrl == "" || pc.populateFile == "" {
		return fmt.Errorf("database url and populate file are required")
	}

	db, err := databases.ParseUrl(pc.databaseUrl)
	if err != nil {
		return fmt.Errorf("failed parsing database url: %w", err)
	}

	ok, err := db.Populate(pc.populateFile)
	if err != nil {
		return fmt.Errorf("failed populating database: %w", err)
	}

	if !ok {
		return fmt.Errorf("failed populating database")
	}

	fmt.Println("Database populated successfully.")

	return nil
}
