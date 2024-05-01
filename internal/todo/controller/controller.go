package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/internal/todo/service"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (tc *TodoController) ListTodoHandler(context *gin.Context) {
	todos := tc.todoService.ListTodo()

	context.JSON(http.StatusOK, todos)
}

func (tc *TodoController) CreateTodoHandler(context *gin.Context) {
	var todo domain.Todo

	bindErr := context.ShouldBind(&todo)

	if bindErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
			"error":   bindErr.Error(),
		})
		return
	}

	savedTodo, insertErr := tc.todoService.InserTodo(&todo)

	if insertErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": insertErr.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, savedTodo)
}

func (tc *TodoController) UpdateTodoHandler(context *gin.Context) {
	id := context.Param("id")

	var todo domain.Todo
	bindErr := context.ShouldBind(&todo)

	if bindErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
			"error":   bindErr.Error(),
		})
		return
	}

	todo.ID = id

	_, updateErr := tc.todoService.UpdateTodo(id, &todo)

	if updateErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": updateErr.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Updated todo",
	})
}

func (tc *TodoController) DeleteTodoHandler(context *gin.Context) {
	id := context.Param("id")

	err := tc.todoService.DeleteTodo(id)

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
