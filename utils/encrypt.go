package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Md5ToLower md5大写
func Md5ToLower(str string) string {
	md5str := Md5String(str)
	md5str = strings.ToLower(md5str)
	return md5str
}

// Md5ToUpper md5小写
func Md5ToUpper(str string) string {
	md5str := Md5String(str)
	md5str = strings.ToUpper(md5str)
	return md5str
}

func Md5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
