package main

import (
	"log"
	"net/http"

	"github.com/gsantosc18/todo/cmd/todo/router"
)

func main() {
	log.Println("Started web server")
	http.ListenAndServe(":8080", router.Initialize())
}
