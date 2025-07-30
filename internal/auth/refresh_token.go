package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func MakeRefreshToken() (string, error) {
	token_data := make([]byte, 32)
	_, err := rand.Read(token_data)
	if err != nil {
		return "", fmt.Errorf("error creating refresh token")
	}

	return hex.EncodeToString(token_data), nil
}
