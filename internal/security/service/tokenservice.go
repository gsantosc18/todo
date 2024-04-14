package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gsantosc18/todo/internal/security/domain"
)

func NewToken(email string) (string, error) {
	user := domain.User{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, user)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ValidateToken(accessToken string) bool {
	_, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid toke: %v", accessToken)
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return err == nil
}
