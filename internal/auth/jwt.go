package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("your_secret_key") // .env-ден алған дұрыс

func GenerateJWT(email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(jwtKey)
}
