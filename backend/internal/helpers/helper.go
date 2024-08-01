package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomString(length int) (string, error) {
	// Generate random bytes
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Encode bytes to a URL-safe base64 string
	randomString := base64.URLEncoding.EncodeToString(bytes)

	// Trim the string to the requested length
	if len(randomString) > length {
		randomString = randomString[:length]
	}

	return randomString, nil
}
