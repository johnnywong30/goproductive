package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateApiKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate api key: %w", err)
	}
	// Encode the key to a base64 string
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}