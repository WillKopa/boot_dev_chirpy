package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed_pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error hashing password: %s", err)
		return "", fmt.Errorf("error hashing password")
	}
	return string(hashed_pw), nil
}

func CheckPasswordHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

}

func GetAPIKey(headers http.Header) (string, error) {
	return GetAuthFromHeader(headers, "ApiKey ")
}

func GetAuthFromHeader(headers http.Header, prefix string) (string, error) {
	auth_header := headers.Get("Authorization")
	if auth_header == "" {
		return "", errors.New("no token in authorization header")
	}

	if !strings.HasPrefix(auth_header, prefix) {
		return "", fmt.Errorf("auth token does not have prefix '%s'", prefix)
	}

	auth_string := strings.TrimPrefix(auth_header, prefix)
	return auth_string, nil
}