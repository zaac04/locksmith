package utilities

import (
	"fmt"
	"os"
)

func LogIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
