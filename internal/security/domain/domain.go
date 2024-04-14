package domain

import "github.com/golang-jwt/jwt"

type User struct {
	Email string
	jwt.StandardClaims
}
