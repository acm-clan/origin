package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func DumpString(s string) string {
	val := ""
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\r':
			val += "\\r"
		case '\n':
			val += "\\n"
		default:
			val += string(s[i])
		}
	}
	return val
}

func DumpStringToByte(s string) string {
	val := ""
	for i := 0; i < len(s); i++ {
		v := int(s[i])
		val += strconv.Itoa(v) + " "
	}
	return val
}
