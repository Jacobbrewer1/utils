package utils

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(data []byte) string {
	hasher := sha256.New()
	hasher.Write(data)
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
