package utilities

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"time"
)

func B64encode(in []byte) (out string) {
	return base64.StdEncoding.EncodeToString(in)
}

func B64Decode(in string) (out []byte, err error) {
	out, err = base64.StdEncoding.DecodeString(in)
	return
}

func GetTimeStamp() string {
	return time.Now().Format(time.RFC850)
}

func GetMD5Hash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func GetCipherName(name string) string {
	return name + ".locksmith"
}
