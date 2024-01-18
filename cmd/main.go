package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sha1sof/Effective_Mobile_test/internal/config"
	"github.com/sha1sof/Effective_Mobile_test/internal/db"
	"github.com/sha1sof/Effective_Mobile_test/internal/routes"
)

func main() {
	err := godotenv.Load("./internal/config/.env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.Init()
	db.InitDB()

	router := mux.NewRouter()

	routes.InitRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
