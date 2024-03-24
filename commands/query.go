package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/luancgs/db-utils/databases"
)

func NewQueryCommand() *QueryCommand {
	queryCommand := &QueryCommand{
		flagSet: flag.NewFlagSet("query", flag.ContinueOnError),
	}

	queryCommand.flagSet.StringVar(&queryCommand.databaseUrl, "url", "", "url of the database")
	queryCommand.flagSet.StringVar(&queryCommand.outputFile, "output", "", "output file")

	return queryCommand
}

type QueryCommand struct {
	flagSet     *flag.FlagSet
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

func (qc *QueryCommand) Run() error {
	if qc.databaseUrl == "" {
		return fmt.Errorf("database url is required")
	}

	qc.query = strings.Join(qc.flagSet.Args(), "\n")

	db, err := databases.ParseUrl(qc.databaseUrl)
	if err != nil {
		return fmt.Errorf("failed parsing database url: %w", err)
	}

	result, err := db.RunQuery(qc.query)
	if err != nil {
		return fmt.Errorf("failed executing query: %w", err)
	}

	if qc.outputFile != "" {
		err := saveResult(qc.outputFile, result)
		if err != nil {
			return fmt.Errorf("failed saving result: %w", err)
		}

		fmt.Println("Query executed successfully.\nResult saved at: ", qc.outputFile)
	} else {
		fmt.Println("Query executed successfully.\nResult: ", result)
	}

	return nil
}

func saveResult(outputFile string, result string) error {
	err := os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return err
	}

	return nil
}
