package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var todo domain.Todo = domain.Todo{
	ID:          "1",
	Name:        "test",
	Description: "Test description",
	Done:        true,
}

func getTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	return ctx
}

func getTestGinContextPost(w *httptest.ResponseRecorder, body any) *gin.Context {
	ctx := getTestGinContext(w)

	data, _ := json.Marshal(body)
	r := bytes.NewBuffer(data)

	ctx.Request.Method = "POST"
	ctx.Request.Body = io.NopCloser(r)

	return ctx
}

func responseMapperTodo[T any](body *bytes.Buffer) T {
	data, _ := io.ReadAll(body)
	var responseTodo T
	json.Unmarshal(data, &responseTodo)
	return responseTodo
}

func TestListTodoSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)

	service := mock.NewMockTodoService(ctrl)

	service.EXPECT().ListTodo(gomock.Any()).Return(domain.NewPaginatedTodo([]domain.Todo{todo}, 0))

	controller := NewTodoController(service)
	controller.ListTodoHandler(ctx)

	listTodo := responseMapperTodo[domain.PaginatedTodo](w.Body)
	responseTodo := listTodo.Data[0]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, todo.ID, responseTodo.ID, "ID expected %s, but got %s", todo.ID, responseTodo.ID)
	assert.Equal(t, todo.Name, responseTodo.Name, "Name expected %s, but got %s", todo.Name, responseTodo.Name)
	assert.Equal(t, todo.Description, responseTodo.Description, "Description expected %s, but got %s", todo.Description, responseTodo.Description)
	assert.Equal(t, todo.Done, responseTodo.Done, "Description expected %v, but got %v", todo.Done, responseTodo.Done)
}

func TestInsertTodoSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContextPost(w, todo)

	service := mock.NewMockTodoService(ctrl)

	service.EXPECT().InserTodo(gomock.Any()).Return(todo, nil)

	controller := NewTodoController(service)

	controller.CreateTodoHandler(ctx)

	responseTodo := responseMapperTodo[domain.Todo](w.Body)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, todo.ID, responseTodo.ID, "ID expected %s, but got %s", todo.ID, responseTodo.ID)
	assert.Equal(t, todo.Name, responseTodo.Name, "Name expected %s, but got %s", todo.Name, responseTodo.Name)
	assert.Equal(t, todo.Description, responseTodo.Description, "Description expected %s, but got %s", todo.Description, responseTodo.Description)
	assert.Equal(t, todo.Done, responseTodo.Done, "Description expected %v, but got %v", todo.Done, responseTodo.Done)
}

func TestInsertTodoErrorBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContextPost(w, []domain.Todo{})

	service := mock.NewMockTodoService(ctrl)

	service.EXPECT().InserTodo(gomock.Any()).Times(0)

	controller := NewTodoController(service)

	controller.CreateTodoHandler(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInsertTodoErrorDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContextPost(w, todo)

	service := mock.NewMockTodoService(ctrl)

	service.
		EXPECT().
		InserTodo(gomock.Any()).
		Return(todo, errors.New("Internal error")).
		Times(1)

	controller := NewTodoController(service)

	controller.CreateTodoHandler(ctx)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateTodoSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContextPost(w, todo)

	expected, _ := json.Marshal(todo)

	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: todo.ID})
	ctx.Request.Body = io.NopCloser(bytes.NewReader(expected))

	service := mock.NewMockTodoService(ctrl)

	service.EXPECT().UpdateTodo("1", gomock.Any()).Return(todo, nil)

	controller := NewTodoController(service)

	controller.UpdateTodoHandler(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTodoErrorBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContextPost(w, []domain.Todo{})

	service := mock.NewMockTodoService(ctrl)

	service.EXPECT().UpdateTodo(gomock.Any(), gomock.Any()).Times(0)

	controller := NewTodoController(service)

	controller.UpdateTodoHandler(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateTodoErrorDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContextPost(w, todo)

	service := mock.NewMockTodoService(ctrl)

	service.EXPECT().
		UpdateTodo(gomock.Any(), gomock.Any()).
		Return(todo, errors.New("internal error")).
		Times(1)

	controller := NewTodoController(service)

	controller.UpdateTodoHandler(ctx)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteTodoSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContextPost(w, todo)

	service := mock.NewMockTodoService(ctrl)
	controller := NewTodoController(service)

	service.EXPECT().DeleteTodo(gomock.Any()).Return(nil)

	controller.DeleteTodoHandler(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTodoErrorDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	ctx := getTestGinContextPost(w, []domain.Todo{})

	service := mock.NewMockTodoService(ctrl)

	service.EXPECT().
		DeleteTodo(gomock.Any()).
		Return(errors.New("Internal error")).
		Times(1)

	controller := NewTodoController(service)

	controller.DeleteTodoHandler(ctx)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
