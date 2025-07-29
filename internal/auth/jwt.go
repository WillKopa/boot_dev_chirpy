package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string
const (
	TokenTypeAcess TokenType = "chirpy"
)

func MakeJWT(user_id uuid.UUID, token_secret string, expires_in time.Duration) (string, error) {
	new_token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer: string(TokenTypeAcess),
		IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expires_in)),
		Subject: user_id.String(),
	})

	return new_token.SignedString([]byte(token_secret))

}

func ValidateJWT(token_string, token_secret string) (uuid.UUID, error) {
	parsed_token, err := jwt.ParseWithClaims(token_string, &jwt.RegisteredClaims{}, func(parsed_token *jwt.Token) (any, error) {
		return []byte(token_secret), nil
	})

	if err != nil {
		log.Printf("error parsing token: %s", err)
		return uuid.Nil, err
	}
	
	user_id, err := parsed_token.Claims.GetSubject();
	if err != nil {
		log.Printf("error getting user id from token: %s", err)
		return uuid.Nil, err
	}

	issuer, err := parsed_token.Claims.GetIssuer();
	if err != nil {
		log.Printf("error issuer does not exist: %s", err)
		return uuid.Nil, err
	}

	if issuer != string(TokenTypeAcess) {
		return uuid.Nil, errors.New("invalid issuer")
	}


	return uuid.Parse(user_id)
}