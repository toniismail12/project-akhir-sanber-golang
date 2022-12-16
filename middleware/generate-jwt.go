package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "pengembara_coding"

// untuk generate token ketika berhasil login
func GenerateJwt(issuer string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(SecretKey))
}
