package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/todo/config/database"
	"github.com/gsantosc18/todo/internal/todo/repository"
	"github.com/gsantosc18/todo/internal/todo/router"
	"github.com/gsantosc18/todo/internal/todo/service"
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

	gin.SetMode(gin.ReleaseMode)

	slog.Info("Started web server")
	route := gin.Default()

	db := database.GetConnect()

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)

	router.GetTodoRoutes(route, todoService)

	route.Run(":8080")
}
