package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sha1sof/Effective_Mobile_test/internal/config"
	"github.com/sha1sof/Effective_Mobile_test/internal/db"
)

func main() {
	err := godotenv.Load("./internal/config/.env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.Init()
	db.InitDB()

	fmt.Println("Все гучи")
}
