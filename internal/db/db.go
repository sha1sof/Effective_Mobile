package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sha1sof/Effective_Mobile_test/internal/config"
)

var db *sqlx.DB

func InitDB() {
	dataName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBUser,
		config.Get().DBPassword,
		config.Get().DBName)

	var err error
	db, err = sqlx.Connect("postgres", dataName)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
}
