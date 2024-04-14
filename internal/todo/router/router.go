package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	securityController "github.com/gsantosc18/todo/internal/security/controller"
	"github.com/gsantosc18/todo/internal/security/service"

	"github.com/gsantosc18/todo/internal/todo/controller"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := "Bearer "
		authorization := c.GetHeader("Authorization")

		if len(authorization) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := authorization[len(bearer):]

		if len(token) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if isValidToken := service.ValidateToken(token); !isValidToken {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func Initialize() {
	route := gin.Default()

	route.POST("/login", securityController.LoginController)

	todo := route.Group("/todo", auth())
	{
		todo.GET("/", controller.ListTodoHandler)
		todo.POST("/", controller.CreateTodoHandler)
		todo.PUT("/:id", controller.UpdateTodoHandler)
		todo.DELETE("/:id", controller.DeleteTodoHandler)
	}

	route.Run(":8080")
}
