package crypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"locksmith/file"
	"locksmith/utilities"
	"os"
	"strings"
)

func EncryptFile(fileName string, key string) {

	l, err := readFile(fileName, key)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
	err = l.encrypt(fileName)
	if err != nil {
		utilities.LogIfError(err)
		return
	}

}

func Encrypt(l *Lock, filename string) {
	err := l.encrypt(filename)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
}

func (l *Lock) encryptWithAES() (err error) {
	gcm, err := l.getAES()

	if err != nil {
		return
	}

	l.nonce = make([]byte, gcm.NonceSize())
	_, err = rand.Read(l.nonce)

	if err != nil {
		return
	}

	l.cipherByte = gcm.Seal(nil, l.nonce, l.message, nil)
	return
}

func (l *Lock) encrypt(filename string) (err error) {
	err = l.encryptWithAES()
	if err != nil {
		return
	}

	Data, err := l.generateData()
	if err != nil {
		return
	}

	file.WriteFile(utilities.GetCipherName(filename), []byte(Data))
	return
}

func (l *Lock) loadMessage(message []byte) {
	l.message = message
}

func (l *Lock) getAES() (gcm cipher.AEAD, err error) {
	block, err := aes.NewCipher(l.aesKey)
	if err != nil {
		return
	}
	gcm, err = cipher.NewGCM(block)
	return
}

func (l *Lock) generateData() (finalData string, err error) {

	header := header{
		LastModified: utilities.GetTimeStamp(),
		CheckSum:     utilities.GetMD5String(l.message),
		Algorithm:    fmt.Sprint("AES ", l.encryption, "bit encryption"),
	}

	headerByte, _ := json.Marshal(header)
	secret := append(l.nonce, l.cipherByte...)

	headerB64 := utilities.B64encode(headerByte)
	data := utilities.B64encode(secret)

	finalData = strings.TrimSpace(headerB64) + "." + strings.TrimSpace(data)
	finalData = strings.TrimSpace(finalData)
	return
}

func FinDiff(originalFile string, CipherFile string, key string) (l *Lock, change bool, err error) {
	_, decrypted, err := getDecryptedValue(&CipherFile, &key)
	if err != nil {
		utilities.LogIfError(err)
		return
	}

	lock, err := readFile(originalFile, key)
	change = findDiff(&decrypted, &lock.message)
	return &lock, change, err
}

func findDiff(cipher *[]byte, rawData *[]byte) (changed bool) {

	cipherBytes := bytes.Split(*cipher, []byte("\n"))
	originalBytes := bytes.Split(*rawData, []byte("\n"))

	line := 0

	var change bytes.Buffer

	for line = 0; line <= len(cipherBytes)-1 && line <= len(originalBytes)-1; line++ {
		compare := bytes.Compare(utilities.GetMD5Hash(cipherBytes[line]), utilities.GetMD5Hash(originalBytes[line]))
		if compare != 0 {
			change.WriteString(fmt.Sprint("\nchange line no:", line+1, "\n"))
			change.WriteString("cipher file: \"")
			change.Write(cipherBytes[line])
			change.WriteString("\"\n")
			change.WriteString("source file: \"")
			change.Write(originalBytes[line])
			change.WriteString("\"\n")
		}
	}

	for i := line; i <= len(cipherBytes)-1; i++ {
		change.WriteString(fmt.Sprint("\nRemoved line:", i+1))
		change.WriteString("\nvalue: \"")
		change.Write(cipherBytes[i])
		change.WriteString("\"\n")
	}

	for i := line; i <= len(originalBytes)-1; i++ {
		change.WriteString(fmt.Sprint("\nAdded line:", i+1))
		change.WriteString("\nvalue: \"")
		change.Write(originalBytes[i])
		change.WriteString("\"\n")
	}
	if change.Len() != 0 {
		fmt.Print("----------")
		fmt.Print(change.String())
		fmt.Println("----------")
		return true
	}
	return false
}

func readFile(filename string, key string) (l Lock, err error) {
	l, err = LoadKey(key)

	if err != nil {
		os.Exit(1)
	}
	data, err := file.ReadFile(filename)

	if err != nil {
		os.Exit(1)
	}

	l.loadMessage(data)
	return
}
