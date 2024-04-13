package repository

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/gsantosc18/todo/config/database"
	"github.com/gsantosc18/todo/internal/todo/domain"
)

type TodoRepositoryImpl struct{}

func NewTodoRepository() *TodoRepositoryImpl {
	return &TodoRepositoryImpl{}
}

func (tri *TodoRepositoryImpl) List() []domain.Todo {
	db := database.GetConnect()

	var todo []domain.Todo
	db.Find(&todo)

	return todo
}

func (tri *TodoRepositoryImpl) Insert(todo *domain.Todo) (domain.Todo, error) {
	db := database.GetConnect()
	todo.ID = uuid.New().String()

	db.Create(&todo)

	slog.Info("Inser new todo", "todo", *todo)

	return *todo, nil
}

func (tri *TodoRepositoryImpl) Update(id string, todo *domain.Todo) (domain.Todo, error) {
	db := database.GetConnect()
	todo.ID = id

	db.Save(&todo)

	return *todo, nil
}

func (tri *TodoRepositoryImpl) Delete(id string) error {
	db := database.GetConnect()

	err := db.Delete(&domain.Todo{}, id)

	return err.Error
}
