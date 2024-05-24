package crypter

type Lock struct {
	b64Key     string
	nonce      []byte
	cipherByte []byte
	aesKey     []byte
	message    []byte
	encryption int
}

type header struct {
	LastModified string
	CheckSum     string
	Algorithm    string
}
