package repository

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/gsantosc18/todo/internal/todo/domain"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db           *gorm.DB
	limitPerPage int
}

func NewTodoRepository(db *gorm.DB, limitPerPage int) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		db:           db,
		limitPerPage: limitPerPage,
	}
}

func (tri *TodoRepositoryImpl) List(page int) *domain.PaginatedTodo {
	var (
		todo  []domain.Todo
		count int64
	)
	var searchPage int

	if page <= 0 {
		searchPage = 0
	} else {
		searchPage = page - 1
	}

	tri.db.Limit(tri.limitPerPage).Offset(tri.limitPerPage * searchPage).Find(&todo)
	tri.db.Model(&domain.Todo{}).Count(&count)
	return domain.NewPaginatedTodo(todo, page, count)
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
