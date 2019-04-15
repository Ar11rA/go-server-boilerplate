package config

import (
	"database/sql"
	"go/go-server-boilerplate/handlers"
	"go/go-server-boilerplate/repository"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	// postgres connection
	_ "github.com/lib/pq"
)

// Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging ...
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Server object
type Server struct {
	Router      *mux.Router
	db          *sql.DB
	userHandler handlers.UserHandler
}

// InitializeDB - returns the db config
func (s *Server) InitializeDB() {
	connectionString := "dbname=temp sslmode=disable"

	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	s.db = db
}

// InitializeEntities - entity config
func (s *Server) InitializeEntities() {
	userRepo := repository.NewUserRepository(s.db)
	userHandler := handlers.NewUserHandler(userRepo)
	s.userHandler = userHandler
}

// InitializeRoutes - add api Routes
func (s *Server) InitializeRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/users", Chain(s.userHandler.GetUsers, Logging())).Methods("GET")
	router.HandleFunc("/users", Chain(s.userHandler.CreateUser, Logging())).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", s.userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", s.userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", s.userHandler.DeleteUser).Methods("DELETE")
	s.Router = router
}

// Run server
func (s *Server) Run() {
	log.Println("Starting server at 8081")
	log.Fatal(http.ListenAndServe(":8081", s.Router))
}
