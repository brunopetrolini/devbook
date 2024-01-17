package authentication

import (
	"devbook/src/infra/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken generates a new access token
func GenerateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey)
}
