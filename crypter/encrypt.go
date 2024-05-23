package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"locksmith/file"
	"locksmith/utilities"
	"os"
	"strings"
)

func (l *Lock) EncryptFile(fileName string) (privKey string, PublicKey string) {

	data, err := file.ReadFile(fileName)
	if err != nil {
		utilities.LogIfError(err)
		os.Exit(1)
	}

	key, err := generateAESkey()

	privKey = l.b64privateKey
	PublicKey = l.b64publicKey

	if err != nil {
		utilities.LogIfError(err)
		return
	}

	err = l.encryptWithAES(key, data)
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


func (l *Lock) encryptWithAES(key []byte, data []byte) (err error) {
	gcm, err := getAES(key)

	l.aesKey = key
	l.message = data
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

	l.cipherByte = gcm.Seal(nil, l.nonce, data, nil)
	return
}

func getAES(key []byte) (gcm cipher.AEAD, err error) {
	block, err := aes.NewCipher(key)

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
	secret := append(l.nonce, l.aesKey...)
	encryptKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, l.publicKey, secret, nil)
	headerB64 := utilities.B64encode(headerByte)
	base64AESkey := utilities.B64encode(encryptKey)
	data := utilities.B64encode(l.cipherByte)
	finalData = strings.TrimSpace(headerB64) + "." + strings.TrimSpace(base64AESkey) + "." + strings.TrimSpace(data)
	finalData = strings.TrimSpace(finalData)
	return
}
