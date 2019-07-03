package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

func ValidateJwt(token string, secret string) bool {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error){
		return []byte(secret), nil
	})

	if !tkn.Valid || err != nil {
		return false
	}

	return true 
}