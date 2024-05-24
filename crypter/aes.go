package crypter

import (
	"crypto/rand"
	"encoding/base64"
)

func New(encryption int) (lock Lock, err error) {
	lock.aesKey, err = generateAESkey(encryption)
	lock.encryption = encryption
	lock.toString()
	return
}

func (e *Lock) GetKey() string {
	return e.b64Key
}

func LoadKey(publicKey string) (lock Lock, err error) {
	lock.b64Key = publicKey
	err = lock.b64decodeKey()
	lock.encryption = len(lock.aesKey) * 8
	return
}

func (e *Lock) toString() {
	e.b64encodeKey()
}

// b64 private keys
func (e *Lock) b64encodeKey() {
	e.b64Key = base64.StdEncoding.EncodeToString(e.aesKey)
}

func (e *Lock) b64decodeKey() (err error) {
	e.aesKey, err = base64.StdEncoding.DecodeString(e.b64Key)
	return
}

func generateAESkey(encryption int) ([]byte, error) {
	key := make([]byte, encryption/8)
	_, err := rand.Read(key)
	return key, err
}
