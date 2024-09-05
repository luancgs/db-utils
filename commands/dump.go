package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/luancgs/db-utils/databases"
	"github.com/luancgs/db-utils/errors"
	"github.com/luancgs/db-utils/views/runners"
)

func NewDumpCommand() *DumpCommand {
	dumpCommand := &DumpCommand{
		flagSet: flag.NewFlagSet("dump", flag.ContinueOnError),
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = ""
	}

	dumpCommand.flagSet.BoolVar(&dumpCommand.help, "help", false, "Show help message")
	dumpCommand.flagSet.BoolVar(&dumpCommand.help, "h", false, "Show help message")

	dumpCommand.flagSet.StringVar(&dumpCommand.dumpDir, "dir", homeDir, "Directory to save the dump file")
	dumpCommand.flagSet.StringVar(&dumpCommand.dumpDir, "d", homeDir, "Directory to save the dump file")

	dumpCommand.flagSet.StringVar(&dumpCommand.databaseUrl, "url", "", "Database URL")
	dumpCommand.flagSet.StringVar(&dumpCommand.databaseUrl, "u", "", "Database URL")

	return dumpCommand
}

type DumpCommand struct {
	flagSet     *flag.FlagSet
	help        bool
	databaseUrl string
	dumpDir     string
}

func (dc *DumpCommand) Name() string {
	return dc.flagSet.Name()
}

func (dc *DumpCommand) Init(args []string) error {
	return dc.flagSet.Parse(args)
}

func (dc *DumpCommand) Run() {
	if dc.help {
		fmt.Println(dc.Help())
		os.Exit(0)
	}

	if dc.databaseUrl == "" {
		runners.ResultRunner("Database URL is required", 0)
		return
	}

	db, err := databases.ParseUrl(dc.databaseUrl)
	errors.ErrorHandler("Error while parsing database URL", err)

	var sqlDump string
	doneChan := make(chan bool)
	var dumpError error

	go func() {
		sqlDump, dumpError = db.Dump(false, dc.dumpDir)
		doneChan <- true
	}()

	runners.LoadingRunner("Dumping database...", nil, doneChan)

	errors.ErrorHandler("Error while dumping database", dumpError)

	runners.ResultRunner(fmt.Sprint("Database dumped successfully. File saved at: ", sqlDump), 1)
}

func (dc DumpCommand) Help() string {
	var output string

	output += fmt.Sprintln("Usage: db-utils dump [flags]")
	output += fmt.Sprintln()
	output += fmt.Sprintln("Flags:")
	output += fmt.Sprintln("  -h, --help         Show help message")
	output += fmt.Sprintln("  -u, --url          Database URL")
	output += fmt.Sprintln("  -d, --dir          Directory to save the dump file")

	return output
}
