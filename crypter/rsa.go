package crypter

import (
	"crypto/rand"
	"encoding/base64"
)

func New() (lock Lock, err error) {
	lock.aesKey, err = generateAESkey()
	lock.ToString()
	return
}

func NewWithKey(Lock Lock, err error) {

}


func LoadKey(publicKey string) (lock Lock, err error) {
	lock.b64Key = publicKey
	err = lock.b64decodeKey()
	return
}

func (e *Lock) ToString() string {
	e.b64encodeKey()
	return e.b64Key
}

// b64 private keys
func (e *Lock) b64encodeKey() {
	e.b64Key = base64.StdEncoding.EncodeToString(e.aesKey)
}

func (e *Lock) b64decodeKey() (err error) {
	e.aesKey, err = base64.StdEncoding.DecodeString(e.b64Key)
	return
}

func generateAESkey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}
