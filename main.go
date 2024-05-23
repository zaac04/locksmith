package main

import (
	"fmt"
	"locksmith/crypter"
	"locksmith/utilities"
	"os"
)

func init() {

}

func main() {

	if len(os.Args) == 3 && os.Args[1] == "-d" {

		var privKey string
		fmt.Print("Enter Private Key: ")
		_, err := fmt.Scanln(&privKey)
		if err != nil {
			utilities.LogIfError(err)
			os.Exit(1)
		}
		crypter.DecryptFile(os.Args[2], privKey)
		os.Exit(0)
	}
	rsa, err := crypter.New()
	utilities.LogIfError(err)

	privkey, _ := rsa.EncryptFile(os.Args[1])
	fmt.Println("PrivateKey")
	fmt.Println(privkey)

}
