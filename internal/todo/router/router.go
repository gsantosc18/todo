package router

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/config/keycloak"
	"github.com/gsantosc18/todo/internal/todo/controller"
	"github.com/gsantosc18/todo/internal/todo/service"
	userController "github.com/gsantosc18/todo/internal/user/controller"
	userService "github.com/gsantosc18/todo/internal/user/service"
)

func GetTodoRoutes(route *gin.Engine, s service.TodoService) {
	context := context.Background()

	keycloak := keycloak.NewKeycloakConfig(
		os.Getenv("KEYCLOAK_HOST"),
		os.Getenv("KEYCLOAK_PORT"),
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("CLIENT_REALM"),
		os.Getenv("ADMIN_USERNAME"),
		os.Getenv("ADMIN_PASSWORD"),
		os.Getenv("ADMIN_REALM"),
	)

	us := userService.NewUserService(context, keycloak)
	uc := userController.NewUserController(us)

	route.POST("/login", uc.LoginController)
	route.POST("/register", uc.CreateNewUserController)

	todoController := controller.NewTodoController(s)

	todo := route.Group("/todo", uc.AuthMiddleware())
	{
		todo.GET("/", todoController.ListTodoHandler)
		todo.POST("/", todoController.CreateTodoHandler)
		todo.PUT("/:id", todoController.UpdateTodoHandler)
		todo.DELETE("/:id", todoController.DeleteTodoHandler)
	}
}
