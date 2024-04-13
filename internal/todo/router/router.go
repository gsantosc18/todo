package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/todo/controller"
)

func Initialize() {
	route := gin.Default()

	route.GET("/todo", controller.ListTodoHandler)
	route.POST("/todo", controller.CreateTodoHandler)
	route.PUT("/todo/:id", controller.UpdateTodoHandler)
	route.DELETE("/todo/:id", controller.DeleteTodoHandler)

	route.Run(":8080")
}
