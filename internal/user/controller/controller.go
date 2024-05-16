package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/user/domain"
	"github.com/gsantosc18/todo/internal/user/service"
)

type UserController struct {
	userService service.UserService
}

type ResponseError struct {
	Error string `json:"error" example:"Internal error"`
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// User login
//
//	@Sumary	Login de usuário
//	@Schemes
//	@Description	Gera o token de acesso do usuário
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.UserLogin	true	"Informações do acesso"
//	@Success		200		{object}	domain.Token
//	@Failure		401		{string}	string	"Acesso não authorizado"
//	@Failure		400		{string}	string	"Parâmetro inválidos"
//	@Router			/login [post]
func (uc *UserController) LoginController(context *gin.Context) {
	var userLogin domain.UserLogin

	bindErr := context.ShouldBind(&userLogin)

	if bindErr != nil {
		context.JSON(http.StatusBadRequest, ResponseError{
			bindErr.Error(),
		})
		return
	}

	token, loginErr := uc.userService.Login(userLogin)

	if loginErr != nil {
		context.JSON(http.StatusUnauthorized, ResponseError{loginErr.Error()})
		return
	}

	context.JSON(http.StatusOK, token)
}

// Create new user
//
//	@Sumary	Criar novo usuário
//	@Schemes
//	@Description	Criar novo usuário
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.CreateUser	true	"Informações do usuário"
//	@Success		200		{object}	domain.User
//	@Failure		400		{string}	string	"Parâmetro inválidos"
//	@Router			/register [post]
func (uc *UserController) CreateNewUserController(context *gin.Context) {
	var createUser domain.CreateUser

	bindErr := context.ShouldBind(&createUser)

	if bindErr != nil {
		context.JSON(http.StatusBadRequest, ResponseError{bindErr.Error()})
		return
	}

	user, err := uc.userService.CreateNewUser(createUser)

	if err != nil {
		context.JSON(http.StatusInternalServerError, ResponseError{err.Error()})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func (uc *UserController) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := "Bearer "
		authorization := strings.TrimPrefix(c.GetHeader("Authorization"), bearer)

		if len(authorization) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := authorization[len(bearer):]

		if len(token) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		_, err := uc.userService.DecodeToken(authorization)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}
	}
}
