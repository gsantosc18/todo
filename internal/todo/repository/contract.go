package repository

import (
	"github.com/gsantosc18/todo/internal/todo/domain"
)

type TodoRepository interface {
	List(page int) *domain.PaginatedTodo
	Insert(todo *domain.Todo) (domain.Todo, error)
	Update(id string, todo *domain.Todo) (domain.Todo, error)
	Delete(id string) error
}
