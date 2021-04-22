package utils

import "math/rand"

// GenRandomStr 生产一串随机字符串
func GenRandomStr(length int) (str string) {
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < length; i++ {
		str += string(chars[rand.Intn(62)])
	}
	return
}
