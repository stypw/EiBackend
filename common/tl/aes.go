package tl

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

const keySize = 32

var (
	aesBlock cipher.Block
)

func init() {
	strKey := "d;zjcjk;cvj0pajiejkl;jkl;jcbnbhjfjfj;asddj;scq23yu89238="

	SetAesKey(strKey)
}

func SetAesKey(strKey string) {
	if strKey == "" {
		return
	}
	strBuff := []byte(strKey)
	keyBuff := strBuff
	for {
		if len(keyBuff) >= keySize {
			break
		}
		keyBuff = append(keyBuff, strBuff...)
	}

	block, err := aes.NewCipher(keyBuff[:keySize])
	if err != nil {
		return
	}
	aesBlock = block
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

//AesEncrypt 加密
func aesEncrypt(data []byte) ([]byte, error) {
	blockSize := aesBlock.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	length := len(encryptBytes)
	crypted := make([]byte, length)
	index := 0
	for index < length {
		start := index
		end := index + blockSize
		aesBlock.Encrypt(crypted[start:end], encryptBytes[start:end])
		index += blockSize
	}
	return crypted, nil
}

//AesDecrypt 解密
func aesDecrypt(data []byte) ([]byte, error) {
	length := len(data)
	crypted := make([]byte, length)
	blockSize := aesBlock.BlockSize()
	index := 0
	for index < length {
		start := index
		end := index + blockSize
		aesBlock.Decrypt(crypted[start:end], data[start:end])
		index += blockSize
	}

	//去除填充
	crypted, err := pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

func AesEncrypt(data string) string {
	res, err := aesEncrypt([]byte(data))
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(res)
}

func AesDecrypt(data string) string {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	dataByte, err = aesDecrypt(dataByte)
	if err != nil {
		return ""
	}
	return string(dataByte)
}
