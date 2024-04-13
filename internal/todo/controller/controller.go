package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/internal/todo/service"
)

var todoService service.TodoService = *service.NewTodoService()

func ListTodoHandler(context *gin.Context) {
	todos := todoService.ListTodo()

	context.JSON(http.StatusOK, todos)
}

func CreateTodoHandler(context *gin.Context) {
	var todo domain.Todo

	bindErr := context.ShouldBind(&todo)

	if bindErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
			"error":   bindErr.Error(),
		})
		return
	}

	savedTodo, insertErr := todoService.InserTodo(&todo)

	if insertErr != nil {
		context.JSON(http.StatusCreated, gin.H{
			"message": "Created new todo",
			"error":   insertErr.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, savedTodo)
}

func UpdateTodoHandler(context *gin.Context) {
	id := context.Param("id")

	var todo domain.Todo
	context.ShouldBind(&todo)

	todo.ID = id

	todoService.UpdateTodo(id, &todo)

	context.JSON(http.StatusOK, gin.H{
		"message": "Updated todo",
	})
}

func DeleteTodoHandler(context *gin.Context) {
	id := context.Param("id")

	err := todoService.DeleteTodo(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "There are a error on delete todo",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted with success",
	})
}
