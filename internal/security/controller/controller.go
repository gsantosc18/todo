package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/security/service"
)

type SecurityController struct {
	tokenService service.TokenService
}

type userLogin struct {
	Email    string `json:"email" example:"email"`
	Password string `json:"password" example:"s3cr3t3"`
}

type tokenResponse struct {
	Token string `json:"token" example:"asdfasdfasdf"`
}

func NewSecurityController(tokenService service.TokenService) *SecurityController {
	return &SecurityController{
		tokenService: tokenService,
	}
}

// Login
//
//	@Summary	Login
//	@Schemes
//	@Description	Gerador de token de acesso
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		controller.userLogin	true	"Requisição de login"
//	@Success		200		{object}	controller.tokenResponse
//	@Router			/login [post]
func (s SecurityController) LoginController(c *gin.Context) {
	var user userLogin
	bindErr := c.ShouldBind(&user)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	if user.Email != "admin" && user.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token, tokenError := s.tokenService.NewToken(user.Email)

	if tokenError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": tokenError.Error(),
		})
		return
	}

	response := tokenResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, response)
}
