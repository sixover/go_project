package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//小写MD5
func Md5Encode(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

//大写
func MD5Encode(password string) string {
	return strings.ToUpper(Md5Encode(password))
}

//加密计算
func PassWdCrpyto(passwd, salt string) string {
	return MD5Encode(passwd + salt)
}

//校验
func ValidPassWd(passwd, salt string, password string) bool {
	return MD5Encode(passwd+salt) == password
}
