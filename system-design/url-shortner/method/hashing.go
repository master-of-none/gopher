package method

import (
	"crypto/sha256"
	"fmt"
)

func Hash(s string) string {
	data := []byte(s)
	hash := sha256.Sum256(data)

	return fmt.Sprintf("%x", hash)[:7]
}
