package main

import (
	"log/slog"

	_ "github.com/gsantosc18/todo/internal/todo/config/log"
	"github.com/gsantosc18/todo/internal/todo/router"
)

func main() {
	slog.Info("Started web server")
	router.Initialize()
}
