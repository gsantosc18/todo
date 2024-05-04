package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gsantosc18/todo/internal/todo/config/database"
	"github.com/gsantosc18/todo/internal/todo/repository"
	"github.com/gsantosc18/todo/internal/todo/router"
	"github.com/gsantosc18/todo/internal/todo/service"
	"github.com/joho/godotenv"

	_ "github.com/gsantosc18/todo/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func logConfig() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(jsonLogger)
}

//	@title						Todo list
//	@version					1.0
//	@description				Poc para estudos de GO
//	@host						localhost:8080
//	@BasePath					/
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
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

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	route.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
