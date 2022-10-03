package goextension

import (
	"encoding/base64"
	"encoding/hex"
)

// String 扩展字符串
type String string

// DecodeBase64 解码Base64
func (s String) DecodeBase64() (Bytes, error) {
	return base64.StdEncoding.DecodeString(string(s))
}

// DecodeHex 解码16进制
func (s String) DecodeHex() (Bytes, error) {
	return hex.DecodeString(string(s))
}

// ToBase64 转换为base64
func (s String) ToBase64() String {
	return Bytes(s).ToBase64()
}

// ToHex 转换为16进制
func (s String) ToHex() String {
	return Bytes(s).ToHex()
}

// Raw 原始对象
func (s String) Raw() string {
	return string(s)
}
