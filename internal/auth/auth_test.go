package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password1 := "password"
	password2 := "electricboogaloo"
	hash1, _:= HashPassword(password1)

	tests := []struct{
		name			string
		password		string
		hash 			string
		want_err		bool
	}{
		{
			name: 			"Correct Password",
			password: 		password1,
			hash: 			hash1,
			want_err: 		false,
		},
		{
			name: 			"Inorrect Password",
			password: 		password2,
			hash: 			hash1,
			want_err: 		true,
		},
		{
			name: 			"Empty Password",
			password: 		"",
			hash: 			hash1,
			want_err: 		true,
		},{
			name: 			"Invalid Hash",
			password: 		password1,
			hash: 			"hash",
			want_err: 		true,
		},
	}

	for _, hash_test := range tests {
		t.Run(hash_test.name, func (t *testing.T) {
			err := CheckPasswordHash(hash_test.password, hash_test.hash)
			if (err != nil) && !hash_test.want_err {
				t.Errorf("CheckPasswordHash() error = %v, wanted error: %v", err, hash_test.want_err)
			}
		})
	}
}