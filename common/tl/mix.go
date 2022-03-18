package tl

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

func reserveSlice(buff []byte) {
	l := len(buff)
	if l < 2 {
		return
	}
	l -= 1
	half := l / 2
	for i := 0; i < half; i++ {
		c := buff[l-i]
		buff[l-i] = buff[i]
		buff[i] = c
	}
}
func reverse(buff []byte, step int) {
	l := len(buff)

	if step < 3 {
		step = 3
	}
	start := 0
	for {
		if start >= l {
			break
		}
		end := start + step
		var child []byte
		if end >= l {
			child = buff[start:]
		} else {
			child = buff[start:end]
		}
		reserveSlice(child)
		start += step
	}
}

func rangeRand(min, max int) int {

	m := &big.Int{}
	m = m.SetInt64(int64(max - min))

	v, r := rand.Int(rand.Reader, m)
	if r != nil {
		return 0
	}

	return int(v.Int64()) + min
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MixEncrypt(str string) string {
	src := []byte(str)
	l := len(src)
	r := byte(rangeRand(0, 255))
	for i := 0; i < l; i++ {
		src[i] += r
	}
	src = append(src, r)
	step := rangeRand(0, intMin(len(src), 255))
	reverse(src, step)
	src = append(src, byte(step))
	return base64.StdEncoding.EncodeToString(src)
}
func MixDecrypt(str string) string {
	src, _ := base64.StdEncoding.DecodeString(string(str))
	l := len(src)
	if l < 3 {
		return str
	}
	step := src[l-1]
	if step > 255 || int(step) > l {
		return str
	}
	reverse(src[:l-1], int(step))
	r := src[l-2]

	for i := 0; i < l-2; i++ {
		src[i] -= r
	}
	return string(src[:l-2])
}
