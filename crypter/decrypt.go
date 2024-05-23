package crypter

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"locksmith/file"
	"locksmith/utilities"
	"strings"
)

func DecryptFile(filename string, PrivKey string) {

	content, err := file.ReadFile(filename)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
	Components := strings.Split(string(content), ".")

	if len(Components) != 3 {
		utilities.LogIfError(fmt.Errorf("unable to decode file due to corrupted data"))
		return
	}

	header, sign, data, err := decodeFileContents(Components)
	if err != nil {
		utilities.LogIfError(err)
		return
	}

	lock := LoadPrivKey(PrivKey)
	err = lock.DecryptRsa(sign)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
	original, err := decryptAES(lock.aesKey[12:], lock.aesKey[:12], data)
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

func decodeFileContents(Components []string) (header header, signature []byte, data []byte, err error) {

	headerBytes, err := utilities.B64Decode(Components[0])
	if err != nil {

		return
	}

	signature, err = utilities.B64Decode(Components[1])
	if err != nil {
		return
	}

	data, err = utilities.B64Decode(Components[2])
	if err != nil {
		return
	}

	err = json.Unmarshal(headerBytes, &header)
	return

}

func (l *Lock) DecryptRsa(cipher []byte) (err error) {
	l.aesKey, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, l.privateKey, cipher, nil)
	return
}

func decryptAES(key []byte, nonce []byte, data []byte) (original []byte, err error) {
	block, err := getAES(key)
	if err != nil {
		return
	}
	original, err = block.Open(nil, []byte(nonce), data, nil)
	return
}
