package crypter

import "crypto/rsa"

type Lock struct {
	privateKey         *rsa.PrivateKey
	publicKey          *rsa.PublicKey
	b64privateKey      string
	b64publicKey       string
	pemBytesPrivateKey []byte
	pemBytesPublicKey  []byte
	nonce              []byte
	cipherByte         []byte
	aesKey             []byte
	message            []byte
}

type header struct {
	Modified string
	CheckSum string
}
