package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gsantosc18/todo/internal/todo/domain"
	"github.com/gsantosc18/todo/internal/todo/repository"
	"github.com/gsantosc18/todo/internal/todo/service"
)

func ListTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(service.ListTodo())
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.Response{Message: "Decode body error", Error: err.Error()})
		return
	}

	insertedTodo, insertErr := repository.InserTodo(&todo)

	if insertErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.Response{Message: "There was an error inserting todo", Error: insertErr.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertedTodo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var todo domain.Todo

	decodeErr := json.NewDecoder(r.Body).Decode(&todo)

	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.Response{Message: "Decode todo error", Error: decodeErr.Error()})
		return
	}

	updatedTodo, updateErr := service.UpdateTodo(id, &todo)

	if updateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.Response{Message: "There are an error updating todo", Error: updateErr.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := service.DeleteTodo(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.Response{Message: "Fail to delete todo", Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain.Response{Message: "Todo deleted with success"})
}
