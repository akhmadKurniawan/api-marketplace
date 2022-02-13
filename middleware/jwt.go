package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	UserID uint64
	Name   string
	Role   int
	jwt.StandardClaims
}

// GetExpiryTime for jwt
func GetExpiryTime(exp time.Duration) int64 {
	return time.Now().Add(time.Hour * exp).Unix()
}

// CreateClaims .
func CreateClaims(userID uint64, name string, role int, exp time.Duration) MyCustomClaims {
	return MyCustomClaims{
		userID,
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: GetExpiryTime(exp),
			IssuedAt:  time.Now().Unix(),
		},
	}
}
