package utilities

import "fmt"

func LogIfError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

