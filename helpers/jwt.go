package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/georgifotev1/go-api/messages"
	"github.com/golang-jwt/jwt"
)

var (
	secretKey        = []byte(os.Getenv("JWT_KEY"))
	blaclistedTokens = make(map[string]bool)
)

func CreateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	if blaclistedTokens[tokenString] {
		return nil, fmt.Errorf(messages.ErrAuthenticationFailed)
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(messages.ErrAuthenticationFailed)
		}

		return secretKey, nil
	})
}

func BlacklistToken(tokenString string) {
	blaclistedTokens[tokenString] = true
}
