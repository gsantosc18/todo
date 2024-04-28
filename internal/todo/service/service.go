package service

import (
	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/internal/todo/repository"
)

type TodoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) *TodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (ts *TodoService) ListTodo() []domain.Todo {
	return ts.todoRepository.List()
}

func (ts *TodoService) InserTodo(todo *domain.Todo) (domain.Todo, error) {
	return ts.todoRepository.Insert(todo)
}

func (ts *TodoService) UpdateTodo(id string, todo *domain.Todo) (domain.Todo, error) {
	return ts.todoRepository.Update(id, todo)
}

func (ts *TodoService) DeleteTodo(id string) error {
	return ts.todoRepository.Delete(id)
}
