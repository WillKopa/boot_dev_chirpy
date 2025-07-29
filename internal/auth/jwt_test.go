package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	user_id := uuid.New()
	token_secret := "it's a secret to everybody"
	expires_in := 5 * time.Minute
	result, err := MakeJWT(user_id, token_secret, expires_in)
	if err != nil {
		t.Fatalf("error making jwt token: %s", err)
	}
	result_user_id, err := ValidateJWT(result, token_secret)
	if err != nil {
		t.Fatalf("error decoding jwt token: %s", err)
	}

	if result_user_id != user_id {
		t.Errorf("user id and decoded user id are not the same")
	}
}