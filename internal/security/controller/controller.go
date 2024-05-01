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
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSecurityController(tokenService service.TokenService) *SecurityController {
	return &SecurityController{
		tokenService: tokenService,
	}
}

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

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
