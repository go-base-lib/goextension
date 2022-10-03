package goextension

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

type Bytes []byte

// ToHex 转二进制
func (b Bytes) ToHex() String {
	return String(hex.EncodeToString(b))
}

// ToBase64 转换为base64
func (b Bytes) ToBase64() String {
	return String(base64.StdEncoding.EncodeToString(b))
}

// DecodeBase64 base64解码
func (b Bytes) DecodeBase64() (Bytes, error) {
	return base64.StdEncoding.DecodeString(string(b))
}

// DecodeHex 16进制解码
func (b Bytes) DecodeHex() (Bytes, error) {
	return hex.DecodeString(string(b))
}

// Equal 是否相等
func (b Bytes) Equal(target Bytes) bool {
	return bytes.Equal(b, target)
}

func (b Bytes) Raw() []byte {
	return b
}
