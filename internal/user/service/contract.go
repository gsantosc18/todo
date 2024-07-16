package service

import "github.com/gsantosc18/todo/internal/user/domain"

type UserService interface {
	CreateNewUser(createUser domain.CreateUser) (*domain.User, domain.ResponseError)
	Login(userLogin domain.UserLogin) (*domain.Token, error)
	DecodeToken(accessToken string) (bool, error)
}
