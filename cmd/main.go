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
		log.Fatal("Error loading .env: ", err)
	}

	config.Init()
	db.InitDB()

	router := mux.NewRouter()
	routes.InitRoutes(router)

	addr := ":8080"
	log.Printf("Server is starting on %s", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
