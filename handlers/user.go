package handlers

import (
	"encoding/json"
	"go/go-server-boilerplate/models"
	"go/go-server-boilerplate/repository"
	"go/go-server-boilerplate/utils"
	"net/http"
)

// GetUsers : get a list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := repository.GetAllUsers()
	utils.JSON(w, http.StatusOK, users)
}

// CreateUser : Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.Error(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()

	createdID := repository.CreateUser(user)
	resp := map[string]interface{}{"status": "User created", "id": createdID}
	utils.JSON(w, http.StatusCreated, resp)
}
