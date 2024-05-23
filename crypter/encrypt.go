package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"locksmith/file"
	"locksmith/utilities"
	"os"
	"strings"
)

func (l *Lock) EncryptFile(fileName string) (Key string) {

	data, err := file.ReadFile(fileName)
	Key = l.b64Key
	if err != nil {
		utilities.LogIfError(err)
		os.Exit(1)
	}

	if err != nil {
		utilities.LogIfError(err)
		return
	}

	err = l.encryptWithAES(data)
	if err != nil {
		utilities.LogIfError(err)
		return
	}

	Data, err := l.generateData()
	if err != nil {
		utilities.LogIfError(err)
		return
	}

	file.WriteFile(utilities.GetCipherName(fileName), []byte(Data))
	return
}

func (l *Lock) encryptWithAES(message []byte) (err error) {
	gcm, err := l.getAES()
	l.loadMessage(message)

	if err != nil {
		utilities.LogIfError(err)
		return
	}

	l.nonce = make([]byte, gcm.NonceSize())
	_, err = rand.Read(l.nonce)

	if err != nil {
		utilities.LogIfError(err)
		return
	}

	l.cipherByte = gcm.Seal(nil, l.nonce, l.message, nil)
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
		Modified: utilities.GetTimeStamp(),
		CheckSum: utilities.GetMD5Hash(l.message),
	}

	headerByte, _ := json.Marshal(header)
	secret := append(l.nonce, l.cipherByte...)

	headerB64 := utilities.B64encode(headerByte)
	data := utilities.B64encode(secret)

	finalData = strings.TrimSpace(headerB64) + "." + strings.TrimSpace(data)
	finalData = strings.TrimSpace(finalData)
	return
}
