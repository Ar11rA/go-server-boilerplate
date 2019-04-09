package handlers

import (
	"encoding/json"
	"go/go-server-boilerplate/repository"
	"net/http"
)

// GetUsers : get a list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := repository.GetAllUsers()
	userBytes, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userBytes)
}
