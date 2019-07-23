package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"obscura-users-backend/db"
)

type Claims struct {
	db.User
	jwt.StandardClaims
}

// CreateNewUserJwt function to create a new jwt token for user information
func CreateNewUserJwt(user db.User, secret string) (string, error) {
	expirationTime := time.Now().Add(60 * time.Hour)
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}