package db

import (
	"time"
)

const (
	base         int64 = 62
	characterSet       = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func toBase62(num int64) string {
	encoded := ""
	for num > 0 {
		r := num % base
		num /= base
		encoded = string(characterSet[r]) + encoded

	}
	return encoded
}

// func ConvertToBase64(input string) string {
// 	return base64.RawStdEncoding.EncodeToString([]byte(input))
// }

// GenerateRandom generates random string based on current datetime
func GenerateRandom() string {
	return toBase62(time.Now().UnixMicro())
}
