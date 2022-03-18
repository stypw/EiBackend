package tl

import (
	"crypto/sha512"
	"encoding/base64"
)

func ToSha512String(input string) string {
	if input == "" {
		return ""
	}
	var e = sha512.Sum512_256([]byte(input))
	return base64.URLEncoding.EncodeToString(e[:])
}

func Encrypt(str string) string {
	return MixEncrypt(AesEncrypt(str))
}
func Decrypt(str string) string {
	return AesDecrypt(MixDecrypt(str))
}
