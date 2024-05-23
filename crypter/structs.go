package crypter

type Lock struct {
	b64Key     string
	nonce      []byte
	cipherByte []byte
	aesKey     []byte
	message    []byte
}

type header struct {
	Modified string
	CheckSum string
}
