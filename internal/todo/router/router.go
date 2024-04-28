package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	securityController "github.com/gsantosc18/todo/internal/security/controller"
	securityService "github.com/gsantosc18/todo/internal/security/service"
	"github.com/gsantosc18/todo/internal/todo/controller"
	"github.com/gsantosc18/todo/internal/todo/service"
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

		if isValidToken := securityService.ValidateToken(token); !isValidToken {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func GetTodoRoutes(route *gin.Engine, s *service.TodoService) {
	route.POST("/login", securityController.LoginController)

	todoController := controller.NewTodoService(s)

	todo := route.Group("/todo", auth())
	{
		todo.GET("/", todoController.ListTodoHandler)
		todo.POST("/", todoController.CreateTodoHandler)
		todo.PUT("/:id", todoController.UpdateTodoHandler)
		todo.DELETE("/:id", todoController.DeleteTodoHandler)
	}
}
