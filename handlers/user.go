package handlers

import (
	"database/sql"
	"encoding/json"
	"go/go-server-boilerplate/models"
	"go/go-server-boilerplate/repository"
	"go/go-server-boilerplate/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// GetUser : get a list of users
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := repository.GetUser(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.Error(w, http.StatusNoContent, "No such user")
		default:
			utils.Error(w, http.StatusInternalServerError, "Internal server error")
		}
	}
	utils.JSON(w, http.StatusOK, user)
}

// UpdateUser : get a list of users
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user := &models.UserDTO{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.Error(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()
	updateErr := repository.UpdateUser(id, user)
	if updateErr != nil {
		utils.Error(w, http.StatusInternalServerError, "Internal server error")
	}
	utils.JSON(w, http.StatusOK, user)
}

// DeleteUser : delete user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	deleteErr := repository.DeleteUser(id)
	if deleteErr != nil {
		utils.Error(w, http.StatusInternalServerError, "Internal server error")
	}
	resp := map[string]interface{}{"status": "User deleted", "id": id}
	utils.JSON(w, http.StatusOK, resp)
}
