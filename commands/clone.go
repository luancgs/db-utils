package commands

import (
	"flag"
	"fmt"

	"github.com/luancgs/db-utils/databases"
)

func NewCloneCommand() *CloneCommand {
	cloneCommand := &CloneCommand{
		flagSet: flag.NewFlagSet("clone", flag.ContinueOnError),
	}

	cloneCommand.flagSet.StringVar(&cloneCommand.databaseOriginUrl, "origin-url", "", "url of the origin database")
	cloneCommand.flagSet.StringVar(&cloneCommand.databaseTargetUrl, "target-url", "", "url of the target database")

	return cloneCommand
}

type CloneCommand struct {
	flagSet           *flag.FlagSet
	databaseOriginUrl string
	databaseTargetUrl string
}

func (qc *CloneCommand) Name() string {
	return qc.flagSet.Name()
}

func (qc *CloneCommand) Init(args []string) error {
	return qc.flagSet.Parse(args)
}

func (qc *CloneCommand) Run() error {

	if qc.databaseOriginUrl == "" || qc.databaseTargetUrl == "" {
		fmt.Print("Enter the origin database url: ")
		var originUrlString string
		fmt.Scanln(&originUrlString)

		fmt.Print("Enter the target database url: ")
		var targetUrlString string
		fmt.Scanln(&targetUrlString)

		originDb, err := databases.ParseUrl(originUrlString)
		if err != nil {
			return fmt.Errorf("failed parsing origin database url: %w", err)
		}

		targetDb, err := databases.ParseUrl(targetUrlString)
		if err != nil {
			return fmt.Errorf("failed parsing target database url: %w", err)
		}

		sqlDump, err := originDb.Dump(true, false, "")
		if err != nil {
			return fmt.Errorf("failed dumping origin database: %w", err)
		}

		ok, err := targetDb.Populate(sqlDump)
		if err != nil {
			return fmt.Errorf("failed populating target database: %w", err)
		}

		if !ok {
			fmt.Println("Database population failed!")
		}
	}

	fmt.Println("Database cloned successfully.")

	return nil
}
