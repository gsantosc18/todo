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
	connect := database.GetConnect()
	rows, err := connect.Query("select * from todo")

	defer rows.Close()

	if err != nil {
		slog.Error(err.Error())
	}

	todos := []domain.Todo{}

	for rows.Next() {
		var todo domain.Todo
		rowErr := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Done)
		if rowErr != nil {
			slog.Error(rowErr.Error())
			continue
		}
		todos = append(todos, todo)
	}

	return todos
}

func (tri *TodoRepositoryImpl) Insert(todo *domain.Todo) (domain.Todo, error) {
	connect := database.GetConnect()

	todo.ID = uuid.New().String()

	slog.Info("Inser new todo", "todo", *todo)

	_, insertErr := connect.Exec("insert into todo (id, name, description, done) values ($1, $2, $3, $4)", todo.ID, todo.Name, todo.Description, todo.Done)

	if insertErr != nil {
		slog.Error(insertErr.Error())
		return domain.Todo{}, insertErr
	}

	return *todo, nil
}

func (tri *TodoRepositoryImpl) Update(id string, todo *domain.Todo) (domain.Todo, error) {
	connect := database.GetConnect()
	_, err := connect.Exec("update todo set name=$1, description=$2, done=$3 where id=$4", todo.Name, todo.Description, todo.Done, id)

	if err != nil {
		slog.Error(err.Error())
		return domain.Todo{}, err
	}

	todo.ID = id

	return *todo, nil
}

func (tri *TodoRepositoryImpl) Delete(id string) error {
	connect := database.GetConnect()

	_, err := connect.Exec("delete from todo where id=$1", id)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
