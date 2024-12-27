package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte("Jiacheng-Todolist-JWT")

type Claims struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// Sign User Token
func GenerateToken(id int64, userName string) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jiacheng",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

// Check User Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
