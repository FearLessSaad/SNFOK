package tooling

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomToken() (string, error) {
	bytes := make([]byte, 32) // 32 bytes * 2 (hex) = 64 characters
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
