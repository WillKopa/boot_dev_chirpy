package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	result, _ := HashPassword(password)

	err := bcrypt.CompareHashAndPassword([]byte(result), []byte(password)) 
	if err != nil {
		t.Errorf("hashed password is not equal to the given password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		t.Fatalf("error hashing password: %s", err)
	}

	err = CheckPasswordHash(password, string(hash))

	if err != nil {
		t.Errorf("hash checker is not working correctly")
	}
}