package utilities

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
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

func GetMD5String(data []byte) string {
	hash := GetMD5Hash(data)
	return hex.EncodeToString(hash)
}

func GetMD5Hash(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}

func GetCipherName(name string) string {
	if strings.Contains(name, ".locksmith") {
		return name
	}
	return name + ".locksmith"
}
