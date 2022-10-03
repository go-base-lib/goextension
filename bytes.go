package goextension

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

type Bytes []byte

// ToHexStr 转16进制字符串
func (b Bytes) ToHexStr() string {
	return hex.EncodeToString(b)
}

// ToBase64Str 转换为base64字符串
func (b Bytes) ToBase64Str() string {
	return base64.StdEncoding.EncodeToString(b)
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
