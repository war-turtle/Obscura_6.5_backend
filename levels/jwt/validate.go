package jwt

import (
	"log"
	"github.com/dgrijalva/jwt-go"
)

func ValidateJwt(token string, secret string) bool {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error){
		return []byte(secret), nil
	})

	if err != nil {
		log.Println(err)
		return false 
	}

	if !tkn.Valid {
		return false
	}

	return true 
}