package method

import (
	"crypto/sha256"
	"fmt"
)

// ! Hashing
func Hash(s string) string {
	data := []byte(s)
	hash := sha256.Sum256(data)

	return fmt.Sprintf("%x", hash)[:7]
}

// ! Base 62
const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base = int64(len(charset))

func EncodeBase62(num int64) string {
	if num == 0 {
		return string(charset[0])
	}
	var encoded []byte

	for num > 0 {
		rem := num % base
		encoded = append([]byte{charset[rem]}, encoded...)
		num /= base
	}
	return string(encoded)
}
