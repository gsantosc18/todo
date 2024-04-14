package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/security/service"
)

type userLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginController(c *gin.Context) {
	var user userLogin
	bindErr := c.ShouldBind(&user)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	if user.Email != "admin" && user.Password != "admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, tokenError := service.NewToken(user.Email)

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
