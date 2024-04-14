package main

import (
	"log/slog"
	"os"

	"github.com/gsantosc18/todo/internal/todo/router"
	"github.com/joho/godotenv"
)

func logConfig() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(jsonLogger)
}

func main() {
	godotenv.Load()
	logConfig()

	slog.Info("Started web server")

	router.Initialize()
}
