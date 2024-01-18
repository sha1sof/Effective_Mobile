package routes

import (
	"github.com/gorilla/mux"
	"github.com/sha1sof/Effective_Mobile_test/internal/handlers"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	router.HandleFunc("/people/delete/{id}", handlers.DeletePerson).Methods("DELETE")
	router.HandleFunc("/people/create", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/people/update/{id}", handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/filter", handlers.Filter).Methods("GET")
}
