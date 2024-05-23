package crypter

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"locksmith/utilities"
)

const (
	pemPrivateKeyType = "RSA PRIVATE KEY"
	pemPublicKeyType  = "PUBLIC KEY"
)

func New() (lock Lock, err error) {
	lock.privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	lock.publicKey = &lock.privateKey.PublicKey
	lock.ToString()
	return
}

func LoadPubKey(publicKey string) (lock Lock) {
	lock.b64publicKey = publicKey
	err := lock.b64decodePubKey()
	if err != nil {
		return
	}

	err = lock.pemDecodePubKey()
	utilities.LogIfError(err)
	return
}

func LoadPrivKey(privateKey string) (lock Lock) {
	lock.b64privateKey = privateKey
	err := lock.b64decodePrivateKey()
	if err != nil {
		return
	}

	err = lock.pemDecodePrivateKey()
	utilities.LogIfError(err)
	return
}

func (e *Lock) ToString() (privateKey string, publicKey string, err error) {

	err = e.pemEncodePrivateKey()

	if err != nil {
		utilities.LogIfError(err)
		return
	}

	e.b64encodePrivKey()
	err = e.pemEncodePublicKey()

	if err != nil {
		utilities.LogIfError(err)
		return
	}

	e.b64encodePubKey()
	return e.b64privateKey, e.b64publicKey, err
}

func (e *Lock) pemEncodePrivateKey() (err error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(e.privateKey)
	e.pemBytesPrivateKey = pem.EncodeToMemory(&pem.Block{
		Type:  pemPrivateKeyType,
		Bytes: privateKeyBytes,
	})
	return
}

func (e *Lock) pemDecodePrivateKey() (err error) {
	block, _ := pem.Decode(e.pemBytesPrivateKey)
	if block == nil || block.Type != pemPrivateKeyType {
		err = fmt.Errorf("failed to load pem")
		utilities.LogIfError(err)
		return
	}
	e.privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	utilities.LogIfError(err)
	return
}

func (e *Lock) pemEncodePublicKey() (err error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(e.publicKey)
	e.pemBytesPublicKey = pem.EncodeToMemory(&pem.Block{
		Type:  pemPublicKeyType,
		Bytes: publicKeyBytes,
	})
	return
}

func (e *Lock) pemDecodePubKey() (err error) {
	block, _ := pem.Decode(e.pemBytesPublicKey)
	if block == nil || block.Type != pemPublicKeyType {
		err = fmt.Errorf("failed to load pem")
		utilities.LogIfError(err)
		return
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	e.publicKey = pub.(*rsa.PublicKey)
	utilities.LogIfError(err)
	return
}

// b64 private keys
func (e *Lock) b64encodePrivKey() {
	e.b64privateKey = base64.StdEncoding.EncodeToString(e.pemBytesPrivateKey)
}

func (e *Lock) b64decodePrivateKey() (err error) {
	e.pemBytesPrivateKey, err = base64.StdEncoding.DecodeString(e.b64privateKey)
	return
}

// b64 public keys
func (e *Lock) b64encodePubKey() {
	e.b64publicKey = base64.StdEncoding.EncodeToString(e.pemBytesPublicKey)
}

func (e *Lock) b64decodePubKey() (err error) {

	e.pemBytesPublicKey, err = base64.StdEncoding.DecodeString(e.b64publicKey)
	return
}

func generateAESkey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}
