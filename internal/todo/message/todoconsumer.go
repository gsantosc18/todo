package message

import (
	"encoding/json"
	"log/slog"

	"github.com/google/uuid"
	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/internal/todo/service"
)

type TodoConsumer struct {
	todoService service.TodoService
}

func NewTodoConsumer(todoService service.TodoService) *TodoConsumer {
	return &TodoConsumer{
		todoService: todoService,
	}
}

func (t *TodoConsumer) Topic() string {
	return "test-go-topic"
}

func (t *TodoConsumer) Receiver(message []byte) {
	var todo domain.Todo

	err := json.Unmarshal(message, &todo)

	slog.Info("Received new message", "todo", todo)

	if err != nil {
		slog.Error("There are an error on consume message", "message", message)
		return
	}

	if todo.ID == "" {
		todo.ID = uuid.NewString()
	}

	savedTodo, insertErr := t.todoService.InserTodo(&todo)

	if insertErr != nil {
		slog.Error("There are an error on insert message", "todo", todo, "error", insertErr.Error())
		return
	}

	slog.Info("Todo created with success", "todo", savedTodo)
}
