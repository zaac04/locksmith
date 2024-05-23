package file

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(filename string) (content []byte, err error) {

	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		return
	}

	content, err = io.ReadAll(file)
	if err != nil {
		return
	}

	defer file.Close()

	fmt.Println("Read", len(content), "bytes from file: ", filename)
	return
}
