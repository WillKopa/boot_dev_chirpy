package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {
	user_id := uuid.New()
	token_secret := "it's a secret to everybody"
	expires_in := 5 * time.Minute
	result, err := MakeJWT(user_id, token_secret, expires_in)

	if err != nil {
		t.Fatalf("error making jwt token: %s", err)
	}

	tests := []struct {
		name				string
		token_string		string
		token_secret		string
		expected_user_ud	uuid.UUID
		want_err			bool
	}{
		{
			name: 				"Valid token",
			token_string: 		result,
			token_secret: 		token_secret,
			expected_user_ud: 	user_id,
			want_err:  			false,
		},
		{
			name: 				"Invalid token",
			token_string: 		"not valid",
			token_secret: 		token_secret,
			expected_user_ud: 	uuid.Nil,
			want_err:  			true,
		},
		{
			name: 				"Wrong secret",
			token_string: 		result,
			token_secret: 		token_secret + "more secret",
			expected_user_ud: 	uuid.Nil,
			want_err:  			true,
		},
	}

	for _, token_test := range tests {
		t.Run(token_test.name, func(t *testing.T) {
			result_user_id, err := ValidateJWT(token_test.token_string, token_test.token_secret)
			
			if (err != nil) && !token_test.want_err {
				t.Errorf("error decoding jwt token: %s", err)
				return
			}

			if result_user_id != token_test.expected_user_ud {
				t.Errorf("user id and decoded user id are not the same")
			}
		})
	}
}