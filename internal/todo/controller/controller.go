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

type response struct {
	Message string `json:"message" example:"Message"`
	Error   string `json:"error" example:"Error message"`
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

// List todos
//
//	@Summary	Lista os todos
//	@Schemes
//	@Description	Listagem de todos cadatrados
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	[]domain.Todo
//	@Failure		401	{string}	string	"Token inválido"
//	@Router			/todo [get]
func (tc *TodoController) ListTodoHandler(context *gin.Context) {
	todos := tc.todoService.ListTodo()

	context.JSON(http.StatusOK, todos)
}

// Create new todo
//
//	@Sumary	Criar novo todo
//	@Schemes
//	@Description	Criar novo todo
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			request	body		domain.Todo	true	"Payload que será criado"
//	@Success		200		{object}	domain.Todo
//	@Failure		401		{object}	controller.response	"Token inválido"
//	@Failure		400		{object}	controller.response	"Request inválido"
//	@Failure		500		{object}	controller.response	"Erro interno do servidor"
//	@Router			/todo	[post]
func (tc *TodoController) CreateTodoHandler(context *gin.Context) {
	var todo domain.Todo

	bindErr := context.ShouldBind(&todo)

	if bindErr != nil {
		context.JSON(http.StatusBadRequest, response{
			Message: "Invalid parameters",
			Error:   bindErr.Error(),
		})
		return
	}

	savedTodo, insertErr := tc.todoService.InserTodo(&todo)

	if insertErr != nil {
		context.JSON(http.StatusInternalServerError, response{
			Error: insertErr.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, savedTodo)
}

// Update todo
//
//	@Sumary	Atualizar todo
//	@Schemes
//	@Description	Atualizar um todo existente
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string		true	"Identificador do todo"
//	@Param			request	body	domain.Todo	true	"Payload que será atualizado"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	controller.response	"Atualizado com sucesso"
//	@Failure		400	{object}	controller.response	"Payload inválido"
//	@Failure		500	{object}	controller.response	"Erro interno do servidor"
//	@Router			/todo/{id} [put]
func (tc *TodoController) UpdateTodoHandler(context *gin.Context) {
	id := context.Param("id")

	var todo domain.Todo
	bindErr := context.ShouldBind(&todo)

	if bindErr != nil {
		context.JSON(http.StatusBadRequest, response{
			Message: "Invalid parameters",
			Error:   bindErr.Error(),
		})
		return
	}

	todo.ID = id

	_, updateErr := tc.todoService.UpdateTodo(id, &todo)

	if updateErr != nil {
		context.JSON(http.StatusInternalServerError, response{
			Error: updateErr.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response{
		Message: "Updated todo",
	})
}

// Delete todo
//
//	@Sumary	Delete todo
//	@Schemes
//	@Description	Apaga um todo pelo identificado
//	@Tags			todo
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Identificado único do todo"
//	@Security		ApiKeyAuth
//	@Success		200			{object}	controller.response
//	@Failure		500			{object}	controller.response	"Erro interno do servidor"
//	@Router			/todo/{id}	[delete]
func (tc *TodoController) DeleteTodoHandler(context *gin.Context) {
	id := context.Param("id")

	err := tc.todoService.DeleteTodo(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, response{
			Message: "There are a error on delete todo",
			Error:   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response{
		Message: "Todo deleted with success",
	})
}
