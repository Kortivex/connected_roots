package hashing

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"path/filepath"
)

func GenUniqueFileName(originalName string) (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("error generating random bytes: %w", err)
	}

	return hex.EncodeToString(randomBytes) + filepath.Ext(originalName), nil
}
