package router

import (
	"github.com/gorilla/mux"
	"github.com/gsantosc18/todo/cmd/todo/handler"
)

func Initialize() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todo", handler.ListTodoHandler).Methods("GET")
	router.HandleFunc("/todo", handler.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo/{id}", handler.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", handler.DeleteTodoHandler).Methods("DELETE")

	return router
}
