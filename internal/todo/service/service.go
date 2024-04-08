package service

import (
	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/internal/todo/repository"
)

func ListTodo() []domain.Todo {
	return repository.ListTodo()
}

func InserTodo(todo domain.Todo) (*domain.Todo, error) {
	return repository.InserTodo(&todo)
}

func UpdateTodo(id string, todo *domain.Todo) (*domain.Todo, error) {
	return repository.UpdateTodo(id, todo)
}

func DeleteTodo(id string) error {
	return repository.DeleteTodo(id)
}
