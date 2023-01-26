package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jeypc/go-jwt-mux/config"
	"github.com/jeypc/go-jwt-mux/helper"
)

const (
	tokenCookie = "token"
)

// JWTMiddleware is the authorization middleware using jwt token
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(tokenCookie)
		if err != nil {
			if err == http.ErrNoCookie {
				helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized",
				})
				return
			}
		}
		tokenString := cookie.Value

		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorExpired:
				helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized, Token expired!",
				})
				return
			case jwt.ValidationErrorSignatureInvalid:
				helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized",
				})
				return
			default:
				helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized",
				})
				return
			}
		}

		if !token.Valid {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
