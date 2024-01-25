package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	secretKey        = []byte(os.Getenv("JWT_KEY"))
	blaclistedTokens = make(map[string]bool)
)

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) error {
	if blaclistedTokens[tokenString] {
		return fmt.Errorf("invalid token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func blacklistToken(tokenString string) {
	blaclistedTokens[tokenString] = true
}
