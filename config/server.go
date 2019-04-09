package config

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// postgres connection
	_ "github.com/lib/pq"
)

// Server object
type Server struct {
	Router *mux.Router
}

// DB conn
var DB *sql.DB

// Initialize the server
func (s *Server) Initialize(router *mux.Router) {
	connectionString := "dbname=temp sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	s.Router = router
}

// Run server
func (s *Server) Run() {
	log.Println("Starting server at 8081")
	log.Fatal(http.ListenAndServe(":8081", s.Router))
}
