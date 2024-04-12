package database

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetConnect() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Houve um erro ao carregar as variáveis de ambiente, erro=%s", err.Error())
	}

	username := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)

	db, dbErr := sql.Open("postgres", connStr)

	if dbErr != nil {
		slog.Error(dbErr.Error())
	}

	return db
}

func init() {
	err := GetConnect().Ping()

	if err != nil {
		slog.Error(err.Error())
	}

	log.Println("Iniciada a conexão com o banco")
}
