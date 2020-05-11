package utils

import (
	"encoding/base64"
)

//加解密
func Encode(username string) string {
	bytes := []byte(username)
	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded
}

func Decode(str string) (decoded string, err error) {
	decodedByte, err := base64.StdEncoding.DecodeString(str)
	decoded = string(decodedByte)
	return
}
