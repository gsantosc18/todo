package controller

import (
	"fmt"
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

type errorResponse struct {
	Error string `json:"error" example:"Error message"`
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
//	@Param			request	body		controller.userLogin	true	"Informações de login"
//	@Success		200		{string}	string					"Token de acesso"
//	@Failure		401		{object}	controller.errorResponse
//	@Failure		400		{object}	controller.errorResponse
//	@Router			/login [post]
func (s SecurityController) LoginController(c *gin.Context) {
	var user userLogin
	bindErr := c.ShouldBind(&user)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: "Invalid parameters",
		})
		return
	}

	if user.Email != "admin" && user.Password != "admin" {
		c.JSON(http.StatusUnauthorized, errorResponse{
			Error: "Invalid credentials",
		})
		return
	}

	token, tokenError := s.tokenService.NewToken(user.Email)

	if tokenError != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: tokenError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Bearer %s", token))
}
