package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 小写MD5
func Md5Encode(data string) string {
	Md5Encoder := md5.New()
	Md5Encoder.Write([]byte(data))
	res := Md5Encoder.Sum(nil)
	return hex.EncodeToString(res)
}

// MD5Encode 大写MD5
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// EncodePassword 密码加密
func EncodePassword(pwd string, randomNumber string) string {
	return Md5Encode(pwd + randomNumber)
}

// ValidatePassword 验证密码
func ValidatePassword(md5Password string, password string, randomNumber string) bool {
	return Md5Encode(password+randomNumber) == md5Password
}
