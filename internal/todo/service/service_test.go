package service

import (
	"testing"

	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/test/mock"
	"go.uber.org/mock/gomock"
)

var todo = domain.Todo{
	ID:          "1",
	Name:        "Test Name",
	Description: "Description Test",
	Done:        true,
}

func TestInsertTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockTodoRepository(ctrl)

	repository.EXPECT().Insert(gomock.Any()).Return(todo, nil)

	service := NewTodoService(repository)

	result, err := service.InserTodo(&todo)

	if err != nil {
		t.Error("Unexpected error when inser todo", "error", err)
		return
	}

	if todo.ID != result.ID {
		t.Errorf("Unexpected id, got [%s]", result.ID)
	}

	if todo.Name != result.Name {
		t.Errorf("Unexpected name, got [%s]", result.Name)
	}

	if todo.Description != result.Description {
		t.Errorf("Unexpected description, got [%s]", result.Description)
	}

	if todo.Done != result.Done {
		t.Errorf("Unexpected done value, got [%v]", result.Done)
	}
}

func TestListTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockTodoRepository(ctrl)

	repository.EXPECT().List(gomock.Any()).Return(domain.NewPaginatedTodo([]domain.Todo{todo}, 0))

	service := NewTodoService(repository)

	todos := service.ListTodo(0)
	firstTodo := todos.Data[0]

	if len(todos.Data) == 0 {
		t.Error("Size of todos is unexpected", "size", len(todos.Data))
		return
	}

	if todo.ID != firstTodo.ID {
		t.Errorf("Unexpected id, got [%s]", firstTodo.ID)
	}

	if todo.Name != firstTodo.Name {
		t.Errorf("Unexpected name, got [%s]", firstTodo.Name)
	}

	if todo.Description != firstTodo.Description {
		t.Errorf("Unexpected description, got [%s]", firstTodo.Description)
	}

	if todo.Done != firstTodo.Done {
		t.Errorf("Unexpected done value, got [%v]", firstTodo.Done)
	}
}

func TestUpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockTodoRepository(ctrl)

	repository.EXPECT().Update("1", &todo).Return(todo, nil)

	service := NewTodoService(repository)

	result, err := service.UpdateTodo("1", &todo)

	if err != nil {
		t.Error("Unexpected error when update todo", "error", err)
		return
	}

	if todo.ID != result.ID {
		t.Errorf("Unexpected id, got [%s]", result.ID)
	}

	if todo.Name != result.Name {
		t.Errorf("Unexpected name, got [%s]", result.Name)
	}

	if todo.Description != result.Description {
		t.Errorf("Unexpected description, got [%s]", result.Description)
	}

	if todo.Done != result.Done {
		t.Errorf("Unexpected done value, got [%v]", result.Done)
	}
}

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockTodoRepository(ctrl)

	repository.EXPECT().Delete("1").Return(nil)

	service := NewTodoService(repository)

	err := service.DeleteTodo("1")

	if err != nil {
		t.Error("Unexpected erro when delete todo", "err", err)
	}
}
