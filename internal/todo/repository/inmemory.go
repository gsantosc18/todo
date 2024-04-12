package repository

import (
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/gsantosc18/todo/internal/todo/domain"
)

type TodoInMemory struct {
	todos []domain.Todo
}

func NewTodoInMemory() *TodoInMemory {
	return &TodoInMemory{todos: []domain.Todo{}}
}

func (t *TodoInMemory) Insert(todo *domain.Todo) (domain.Todo, error) {
	if todo == nil {
		slog.Error("Insert failed because todo is null")
		return domain.Todo{}, errors.New("Todo is null")
	}

	todo.ID = uuid.NewString()

	t.todos = append(t.todos, *todo)
	slog.Info("Inserted new todo with success", "todo", *todo)
	return *todo, nil
}

func (t *TodoInMemory) List() []domain.Todo {
	return t.todos
}

func (t *TodoInMemory) findById(id string) (*domain.Todo, error) {
	var todo *domain.Todo

	for i := range t.todos {
		td := &t.todos[i]
		if td.ID == id {
			todo = td
			break
		}
	}

	if todo == nil {
		return nil, errors.New("Not found todo by id")
	}

	return todo, nil
}

func (t *TodoInMemory) Update(id string, todo *domain.Todo) (domain.Todo, error) {
	if todo == nil {
		slog.Error("Update todo failed because todo is required", "id", id)
		return domain.Todo{}, errors.New("Todo is required")
	}

	savedTodo, err := t.findById(id)

	if err != nil {
		slog.Error("Update todo failed", "todo", *todo, "error", err.Error())
		return domain.Todo{}, err
	}

	savedTodo.Name = todo.Name
	savedTodo.Description = todo.Description
	savedTodo.Done = todo.Done

	slog.Info("Updated todo with success", "todo", *todo)

	return *savedTodo, nil
}

func (t *TodoInMemory) Delete(id string) error {
	var position int = -1

	for i := range t.todos {
		td := &t.todos[i]
		if td.ID == id {
			position = i
			break
		}
	}

	if position == -1 {
		return errors.New("Not found todo by id")
	}

	t.todos = append(t.todos[:position], t.todos[position+1:]...)

	slog.Info("Todo was deleted with success", "id", id)

	return nil
}
