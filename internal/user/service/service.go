package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gsantosc18/todo/internal/config/keycloak"
	"github.com/gsantosc18/todo/internal/user/domain"
)

type UserServiceImpl struct {
	ctx            context.Context
	keycloakConfig *keycloak.KeycloakConfig
}

func NewUserService(ctx context.Context, keycloakConfig *keycloak.KeycloakConfig) *UserServiceImpl {
	return &UserServiceImpl{
		ctx:            ctx,
		keycloakConfig: keycloakConfig,
	}
}

func (usi *UserServiceImpl) getClientKeycloak() *gocloak.GoCloak {
	keycloak_host := fmt.Sprintf("http://%s:%s", usi.keycloakConfig.Server.Host, usi.keycloakConfig.Server.Port)
	return gocloak.NewClient(keycloak_host)
}

func (usi *UserServiceImpl) CreateNewUser(createUser domain.CreateUser) (*domain.User, domain.ResponseError) {
	keycloakUser := gocloak.User{
		FirstName: &createUser.FirstName,
		LastName:  &createUser.LastName,
		Email:     &createUser.Email,
		Enabled:   gocloak.BoolP(true),
		Username:  &createUser.UserName,
	}

	client := usi.getClientKeycloak()
	token, tokenErr := client.LoginAdmin(usi.ctx, usi.keycloakConfig.Admin.Username, usi.keycloakConfig.Admin.Password, usi.keycloakConfig.Admin.Realm)

	if tokenErr != nil {
		return nil, domain.ResponseError{
			Error:      fmt.Errorf("keycloak: [%s]", removeStatus(tokenErr.Error())),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if token == nil {
		return nil, domain.ResponseError{
			Error:      errors.New("invalid client secret"),
			StatusCode: http.StatusInternalServerError,
		}
	}

	userId, createUserErr := client.CreateUser(usi.ctx, token.AccessToken, usi.keycloakConfig.Client.Realm, keycloakUser)

	if createUserErr != nil {
		return nil, domain.ResponseError{
			Error:      fmt.Errorf("keycloak: [%s]", removeStatus(createUserErr.Error())),
			StatusCode: http.StatusBadRequest,
		}
	}

	passwordErr := client.SetPassword(usi.ctx, token.AccessToken, userId, usi.keycloakConfig.Client.Realm, createUser.Password, false)

	if passwordErr != nil {
		return nil, domain.ResponseError{
			Error:      passwordErr,
			StatusCode: http.StatusBadRequest,
		}
	}

	user := &domain.User{
		ID:        userId,
		FirstName: createUser.FirstName,
		LastName:  createUser.LastName,
		UserName:  createUser.UserName,
		Email:     createUser.Email,
		Enabled:   true,
	}

	return user, domain.ResponseError{}
}

func (usi *UserServiceImpl) Login(userLogin domain.UserLogin) (*domain.Token, error) {
	client := usi.getClientKeycloak()

	login, err := client.Login(
		usi.ctx, usi.keycloakConfig.Client.ClientID,
		usi.keycloakConfig.Client.ClientSecret,
		usi.keycloakConfig.Client.Realm,
		userLogin.Username,
		userLogin.Password,
	)

	if err != nil {
		return nil, errors.New(removeStatus(err.Error()))
	}

	return &domain.Token{
		Token:     login.AccessToken,
		ExpiredIn: login.ExpiresIn,
	}, nil
}

func (usi *UserServiceImpl) DecodeToken(accessToken string) (bool, error) {
	client := usi.getClientKeycloak()

	jwt, _, err := client.DecodeAccessToken(usi.ctx, accessToken, usi.keycloakConfig.Client.Realm)

	if err != nil {
		return false, errors.New(removeStatus(err.Error()))
	}

	return jwt.Valid, nil
}

func removeStatus(errorMessage string) string {
	return errorMessage[4:]
}
