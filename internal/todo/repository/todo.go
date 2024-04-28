package repository

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/gsantosc18/todo/internal/todo/domain"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		db: db,
	}
}

func (tri *TodoRepositoryImpl) List() []domain.Todo {
	var todo []domain.Todo
	tri.db.Find(&todo)

	return todo
}

func (tri *TodoRepositoryImpl) Insert(todo *domain.Todo) (domain.Todo, error) {
	todo.ID = uuid.New().String()

	tri.db.Create(&todo)

	slog.Info("Inser new todo", "todo", *todo)

	return *todo, nil
}

func (tri *TodoRepositoryImpl) Update(id string, todo *domain.Todo) (domain.Todo, error) {
	todo.ID = id

	tri.db.Save(&todo)

	return *todo, nil
}

func (tri *TodoRepositoryImpl) Delete(id string) error {

	err := tri.db.Delete(&domain.Todo{ID: id})

	return err.Error
}
