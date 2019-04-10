package routes

import (
	"go/go-server-boilerplate/handlers"

	"github.com/gorilla/mux"
)

// Router - Api routes
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	return router
}
