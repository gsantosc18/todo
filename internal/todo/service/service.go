package service

import (
	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/internal/todo/repository"
)

type TodoServiceImpl struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepository: todoRepository}
}

func (ts *TodoServiceImpl) ListTodo(page int) *domain.PaginatedTodo {
	return ts.todoRepository.List(page)
}

func (ts *TodoServiceImpl) InserTodo(todo *domain.Todo) (domain.Todo, error) {
	return ts.todoRepository.Insert(todo)
}

func (ts *TodoServiceImpl) UpdateTodo(id string, todo *domain.Todo) (domain.Todo, error) {
	return ts.todoRepository.Update(id, todo)
}

func (ts *TodoServiceImpl) DeleteTodo(id string) error {
	return ts.todoRepository.Delete(id)
}
