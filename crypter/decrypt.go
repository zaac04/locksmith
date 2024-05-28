package crypter

import (
	"encoding/json"
	"fmt"
	"locksmith/file"
	"locksmith/utilities"
	"strings"
)

func DecryptFile(filename string, key string) {

	header, original, err := GetDecryptedValue(&filename, &key)

	if err != nil {
		utilities.LogIfError(err)
		return
	}

	ok := compareChecksum(header.CheckSum, original)

	if !ok {
		fmt.Println("failed to write file!")
		return
	}

	file.WriteFile(".env", original)
}

func decodeFileContents(Components []string) (header header, data []byte, err error) {

	headerBytes, err := utilities.B64Decode(Components[0])
	if err != nil {
		return
	}

	data, err = utilities.B64Decode(Components[1])
	if err != nil {
		return
	}

	err = json.Unmarshal(headerBytes, &header)
	return

}

func (l *Lock) decryptAES(data *[]byte) (original []byte, err error) {
	block, err := l.getAES()

	nonce := block.NonceSize()
	l.cipherByte = *data

	if err != nil {
		return
	}
	original, err = block.Open(nil, l.cipherByte[:nonce], l.cipherByte[nonce:], nil)
	return
}

func MatchCheckSum(original string, cipher string) {

	content, err := file.ReadFile(original)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
	header, _, err := loadCipherFile(cipher)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
	compareChecksum(header.CheckSum, content)
}

func compareChecksum(val1 string, val2 []byte) bool {
	if val1 != utilities.GetMD5String(val2) {
		fmt.Println("Checksum Failed!")
		return false
	}
	fmt.Println("CheckSum Matched!")
	return true
}

func loadCipherFile(filename string) (header header, data []byte, err error) {
	content, err := file.ReadFile(filename)
	if err != nil {
		return
	}
	Components := strings.Split(string(content), ".")

	if len(Components) != 2 {
		err = fmt.Errorf("unable to decode file due to corrupted data")
		return
	}

	header, data, err = decodeFileContents(Components)
	return
}

func ReadCipherFile(filename string) {
	header, data, err := loadCipherFile(filename)
	if err != nil {
		utilities.LogIfError(err)
		return
	}
	fmt.Println("--------")
	utilities.PrintStruct(header)
	fmt.Println("Cipher:", len(data), "Bytes")
	fmt.Println("--------")
}

func GetDecryptedValue(file *string, key *string) (header header, cipher []byte, err error) {

	header, data, err := loadCipherFile(*file)

	if err != nil {
		return
	}

	lock, err := LoadKey(*key)
	if err != nil {
		return
	}

	cipher, err = lock.decryptAES(&data)
	return
}
