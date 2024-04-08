package repository

import (
	"errors"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/gsantosc18/todo/internal/todo/domain"
)

var todos []domain.Todo = []domain.Todo{}
var log *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func InserTodo(todo *domain.Todo) (*domain.Todo, error) {
	if todo == nil {
		log.Error("Insert failed because todo is null")
		return nil, errors.New("Todo is null")
	}

	todo.ID = uuid.NewString()

	todos = append(todos, *todo)
	log.Info("Inserted new todo with success", "todo", *todo)
	return todo, nil
}

func ListTodo() []domain.Todo {
	return todos
}

func FindById(id string) (*domain.Todo, error) {
	var todo *domain.Todo

	for t := range todos {
		td := &todos[t]
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

func UpdateTodo(id string, todo *domain.Todo) (*domain.Todo, error) {
	if todo == nil {
		log.Error("Update todo failed because todo is required", "id", id)
		return nil, errors.New("Todo is required")
	}

	savedTodo, err := FindById(id)

	if err != nil {
		log.Error("Update todo failed", "todo", *todo, "error", err.Error())
		return nil, err
	}

	savedTodo.Name = todo.Name
	savedTodo.Description = todo.Description
	savedTodo.Done = todo.Done

	log.Info("Updated todo with success", "todo", *todo)

	return savedTodo, nil
}

func DeleteTodo(id string) error {
	var position int = -1

	for i := range todos {
		td := todos[i]
		if td.ID == id {
			position = i
			break
		}
	}

	if position == -1 {
		return errors.New("Not found todo by id")
	}

	todos = append(todos[:position], todos[position+1:]...)

	log.Info("Todo was deleted with success", "id", id)

	return nil
}
