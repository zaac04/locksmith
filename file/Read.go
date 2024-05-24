package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (content []byte, err error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		return
	}
	defer file.Close()

	file_info, _ := file.Stat()
	content = make([]byte, file_info.Size())
	reader := bufio.NewReader(file)
	reader.Read(content)

	fmt.Println("Read", len(content), "bytes from file: ", filename)
	return
}
