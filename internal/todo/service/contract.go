package service

import "github.com/gsantosc18/todo/internal/todo/domain"

type TodoService interface {
	ListTodo() []domain.Todo
	InserTodo(todo *domain.Todo) (domain.Todo, error)
	UpdateTodo(id string, todo *domain.Todo) (domain.Todo, error)
	DeleteTodo(id string) error
}
