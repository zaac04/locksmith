package file

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(FileName string, Content []byte) error {
	file, err := os.OpenFile(FileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n, err := file.Write(Content)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully wrote", n, " bytes to File : ", FileName)
	return err
}
