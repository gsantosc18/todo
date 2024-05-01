package service

type TokenService interface {
	NewToken(email string) (string, error)
	ValidateToken(accessToken string) bool
}
