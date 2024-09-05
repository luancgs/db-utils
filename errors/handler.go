package errors

import (
	"fmt"
	"os"
)

func ErrorHandler(text string, err error) {
	if err != nil {
		output := fmt.Sprintf("%s: %s", text, err.Error())

		fmt.Println(output)
		os.Exit(1)
	}
}
