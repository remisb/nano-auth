package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")
var tokens []string

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func AddToken(token string) {
	tokens = append(tokens, token)
}

func EncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidateToken(token string) bool {
	for _, v := range tokens {
		if v == token {
			return true
		}
	}
	return false
}

func ParseWithClaims(tokenString string) {

}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func ResetPassword(email string) error {
	if len(email) == 0 {
		return errors.New("auth: Invalid email argument value")
	}

	// sending an email introduces mailing/messaging system dependency

	return errors.New("auth: Not implemented yet")
}
