package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/sha1sof/Effective_Mobile_test/internal/handlers"
)

func InitRoutes(router *mux.Router) {

	log.Println("Initializing routes...")

	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	log.Println("Added route: GET /people")

	router.HandleFunc("/people/delete/{id}", handlers.DeletePerson).Methods("DELETE")
	log.Println("Added route: DELETE /people/delete/{id}")

	router.HandleFunc("/people/create", handlers.CreatePerson).Methods("POST")
	log.Println("Added route: POST /people/create")

	router.HandleFunc("/people/update/{id}", handlers.UpdatePerson).Methods("PUT")
	log.Println("Added route: PUT /people/update/{id}")

	router.HandleFunc("/people/filter", handlers.Filter).Methods("GET")
	log.Println("Added route: GET /people/filter")

	log.Println("Routes initialized successfully.")
}
