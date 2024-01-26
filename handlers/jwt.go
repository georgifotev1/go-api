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

func createToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	if blaclistedTokens[tokenString] {
		return nil, fmt.Errorf("invalid token")
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
}

func blacklistToken(tokenString string) {
	blaclistedTokens[tokenString] = true
}
