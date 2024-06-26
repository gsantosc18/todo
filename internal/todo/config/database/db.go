package database

import (
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type db struct {
	Host     string
	User     string
	Password string
	Dbname   string
}

func getDbInfo() *db {
	return &db{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}
}

func GetConnect() *gorm.DB {
	dbInfo := getDbInfo()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbInfo.User, dbInfo.Password, dbInfo.Dbname, dbInfo.Host)

	db, dbErr := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if dbErr != nil {
		slog.Error(dbErr.Error())
	}

	return db
}
