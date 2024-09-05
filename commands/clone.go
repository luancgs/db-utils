package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/luancgs/db-utils/databases"
	"github.com/luancgs/db-utils/errors"
	"github.com/luancgs/db-utils/views/runners"
)

func NewCloneCommand() *CloneCommand {
	cloneCommand := &CloneCommand{
		flagSet: flag.NewFlagSet("clone", flag.ContinueOnError),
	}

	cloneCommand.flagSet.BoolVar(&cloneCommand.help, "help", false, "Show help message")
	cloneCommand.flagSet.BoolVar(&cloneCommand.help, "h", false, "Show help message")

	cloneCommand.flagSet.StringVar(&cloneCommand.sourceUrl, "input", "", "Source database URL")
	cloneCommand.flagSet.StringVar(&cloneCommand.sourceUrl, "i", "", "Source database URL")

	cloneCommand.flagSet.StringVar(&cloneCommand.targetUrl, "output", "", "Target database URL")
	cloneCommand.flagSet.StringVar(&cloneCommand.targetUrl, "o", "", "Target database URL")

	return cloneCommand
}

type CloneCommand struct {
	flagSet   *flag.FlagSet
	help      bool
	sourceUrl string
	targetUrl string
}

func (qc *CloneCommand) Name() string {
	return qc.flagSet.Name()
}

func (qc *CloneCommand) Init(args []string) error {
	return qc.flagSet.Parse(args)
}

func (qc *CloneCommand) Run() {

	if qc.help {
		fmt.Println(qc.Help())
		os.Exit(0)
	}

	if qc.sourceUrl == "" || qc.targetUrl == "" {
		runners.ResultRunner("Source and Target URLs are required", 0)
		return
	}

	originDb, err := databases.ParseUrl(qc.sourceUrl)
	errors.ErrorHandler("Error while parsing source database URL", err)

	targetDb, err := databases.ParseUrl(qc.targetUrl)
	errors.ErrorHandler("Error while parsing target database URL", err)

	sqlDump, err := originDb.Dump(true, "")
	errors.ErrorHandler("Error while dumping origin database", err)

	ok, err := targetDb.Restore(sqlDump)
	errors.ErrorHandler("Error while restoring target database", err)

	if !ok {
		runners.ResultRunner("Database restore failed!", 0)
		return
	}

	runners.ResultRunner("Database cloned successfully.", 1)
}

func (qc CloneCommand) Help() string {
	var output string

	output += fmt.Sprintln("Usage: db-utils clone [flags]")
	output += fmt.Sprintln()
	output += fmt.Sprintln("Flags:")
	output += fmt.Sprintln("  -h, --help      Show help message")
	output += fmt.Sprintln("  -i, --input     Source database URL")
	output += fmt.Sprintln("  -o, --output    Target database URL")

	return output
}
