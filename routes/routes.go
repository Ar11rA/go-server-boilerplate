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
	return router
}
