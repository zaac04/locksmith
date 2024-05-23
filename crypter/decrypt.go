package crypter

import (
	"encoding/json"
	"fmt"
	"locksmith/file"
	"locksmith/utilities"
	"strings"
)

func DecryptFile(filename string, key string) {

	content, err := file.ReadFile(filename)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
	Components := strings.Split(string(content), ".")

	if len(Components) != 2 {
		utilities.LogIfError(fmt.Errorf("unable to decode file due to corrupted data"))
		return
	}

	header, data, err := decodeFileContents(Components)
	if err != nil {
		utilities.LogIfError(err)
		return
	}

	lock, err := LoadKey(key)
	if err != nil {
		utilities.LogIfError(err)
		return
	}

	original, err := lock.decryptAES(data)
	if err != nil {
		utilities.LogIfError(err)
		return
	}

	if header.CheckSum != utilities.GetMD5Hash(original) {
		fmt.Println("Checksum Failed!")
		return
	}

	fmt.Println("\nChecksum Matches!")
	file.WriteFile(".env", original)

}

func decodeFileContents(Components []string) (header header, data []byte, err error) {

	headerBytes, err := utilities.B64Decode(Components[0])
	if err != nil {
		return
	}

	data, err = utilities.B64Decode(Components[1])
	if err != nil {
		return
	}

	err = json.Unmarshal(headerBytes, &header)
	return

}

func (l *Lock) decryptAES(data []byte) (original []byte, err error) {
	block, err := l.getAES()

	nonce := block.NonceSize()
	l.cipherByte = data

	if err != nil {
		return
	}
	original, err = block.Open(nil, l.cipherByte[:nonce], l.cipherByte[nonce:], nil)
	return
}
