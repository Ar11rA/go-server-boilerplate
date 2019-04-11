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

type userHandler struct {
	repo repository.UserRepository
}

// UserHandler ...
type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

// NewUserHandler ...
func NewUserHandler(repo repository.UserRepository) UserHandler {
	return &userHandler{
		repo: repo,
	}
}

// GetUsers : get a list of users
func (u *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := u.repo.GetAllUsers()
	utils.JSON(w, http.StatusOK, users)
}

// CreateUser : Create new user
func (u *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.Error(w, http.StatusBadRequest, "Bad request")
		return
	}
	defer r.Body.Close()

	createdID := u.repo.CreateUser(user)
	resp := map[string]interface{}{"status": "User created", "id": createdID}
	utils.JSON(w, http.StatusCreated, resp)
}

// GetUser : get a list of users
func (u *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := u.repo.GetUser(id)
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
func (u *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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
	updateErr := u.repo.UpdateUser(id, user)
	if updateErr != nil {
		utils.Error(w, http.StatusInternalServerError, "Internal server error")
	}
	utils.JSON(w, http.StatusOK, user)
}

// DeleteUser : delete user by id
func (u *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	deleteErr := u.repo.DeleteUser(id)
	if deleteErr != nil {
		utils.Error(w, http.StatusInternalServerError, "Internal server error")
	}
	resp := map[string]interface{}{"status": "User deleted", "id": id}
	utils.JSON(w, http.StatusOK, resp)
}
