package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jeypc/go-jwt-mux/config"
)

const (
	issuerName = "davidalexander.com"
)

// GenerateToken will create a new jwt token.
//
// Return token string and nil error when success.
// Otherwise, empty string and non-nil error.
func GenerateToken(username string) (string, error) {
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuerName,
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenAlgo.SignedString(config.JWT_KEY)
}
