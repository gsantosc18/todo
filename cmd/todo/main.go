package main

import (
	"log"
	"net/http"

	_ "github.com/gsantosc18/todo/config/log"
	"github.com/gsantosc18/todo/internal/todo/router"
)

func main() {
	log.Println("Started web server")
	http.ListenAndServe(":8080", router.Initialize())
}
