package router

import (
	"github.com/gorilla/mux"
	"github.com/gsantosc18/todo/internal/todo/controller"
)

func Initialize() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todo", controller.ListTodoHandler).Methods("GET")
	router.HandleFunc("/todo", controller.CreateTodoHandler).Methods("POST")
	router.HandleFunc("/todo/{id}", controller.UpdateTodoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", controller.DeleteTodoHandler).Methods("DELETE")

	return router
}
