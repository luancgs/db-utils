package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/luancgs/db-utils/databases"
	"github.com/luancgs/db-utils/errors"
	"github.com/luancgs/db-utils/views/runners"
)

func NewQueryCommand() *QueryCommand {
	queryCommand := &QueryCommand{
		flagSet: flag.NewFlagSet("query", flag.ContinueOnError),
	}

	queryCommand.flagSet.BoolVar(&queryCommand.help, "help", false, "Show help message")
	queryCommand.flagSet.BoolVar(&queryCommand.help, "h", false, "Show help message")

	queryCommand.flagSet.StringVar(&queryCommand.databaseUrl, "url", "", "Database URL")
	queryCommand.flagSet.StringVar(&queryCommand.databaseUrl, "u", "", "Database URL")

	queryCommand.flagSet.StringVar(&queryCommand.outputFile, "output", "", "Output file for the query")
	queryCommand.flagSet.StringVar(&queryCommand.outputFile, "o", "", "Output file for the query")

	return queryCommand
}

type QueryCommand struct {
	flagSet     *flag.FlagSet
	help        bool
	databaseUrl string
	query       string
	outputFile  string
}

func (qc *QueryCommand) Name() string {
	return qc.flagSet.Name()
}

func (qc *QueryCommand) Init(args []string) error {
	return qc.flagSet.Parse(args)
}

func (qc *QueryCommand) Run() {
	if qc.help {
		fmt.Println(qc.Help())
		os.Exit(0)
	}

	if qc.databaseUrl == "" {
		runners.ResultRunner("Database URL is required", 0)
		return
	}

	qc.query = strings.Join(qc.flagSet.Args(), "\n")

	db, err := databases.ParseUrl(qc.databaseUrl)
	errors.ErrorHandler("Error while parsing database URL", err)

	result, err := db.RunQuery(qc.query)
	errors.ErrorHandler("Error while executing query", err)

	if qc.outputFile != "" {
		err := saveResult(qc.outputFile, result)
		errors.ErrorHandler("Error while saving query result to file", err)

		fmt.Println("Query executed successfully.\nResult saved at: ", qc.outputFile)
	} else {
		fmt.Println("Query executed successfully. Result:\n\n", result)
	}
}

func saveResult(outputFile string, result string) error {
	err := os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (qc QueryCommand) Help() string {
	var output string

	output += fmt.Sprintln("Usage: db-utils dump [flags]")
	output += fmt.Sprintln()
	output += fmt.Sprintln("Flags:")
	output += fmt.Sprintln("  -h, --help      Show help message")
	output += fmt.Sprintln("  -u, --url       Database URL")
	output += fmt.Sprintln("  -o, --output    Output file for the query")

	return output
}
