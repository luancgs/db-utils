package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/luancgs/db-utils/databases"
)

func NewDumpCommand() *DumpCommand {
	dumpCommand := &DumpCommand{
		flagSet: flag.NewFlagSet("dump", flag.ContinueOnError),
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = ""
	}

	dumpCommand.flagSet.StringVar(&dumpCommand.dumpDir, "dir", homeDir, "directory to save the dump file")
	dumpCommand.flagSet.StringVar(&dumpCommand.databaseUrl, "url", "", "url of the database")
	dumpCommand.flagSet.BoolVar(&dumpCommand.isRestorable, "restorable", false, "if the dump should be restorable")

	return dumpCommand
}

type DumpCommand struct {
	flagSet      *flag.FlagSet
	databaseUrl  string
	dumpDir      string
	isRestorable bool
}

func (dc *DumpCommand) Name() string {
	return dc.flagSet.Name()
}

func (dc *DumpCommand) Init(args []string) error {
	return dc.flagSet.Parse(args)
}

func (dc *DumpCommand) Run() error {
	if dc.databaseUrl == "" {
		return fmt.Errorf("database url is required")
	}

	db, err := databases.ParseUrl(dc.databaseUrl)
	if err != nil {
		return fmt.Errorf("failed parsing database url: %w", err)
	}

	sqlDump, err := db.Dump(false, dc.isRestorable, dc.dumpDir)
	if err != nil {
		return fmt.Errorf("failed dumping database: %w", err)
	}

	fmt.Println("Database dumped successfully. File saved at: ", sqlDump)

	return nil
}
