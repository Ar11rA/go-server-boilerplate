package repository

import (
	"database/sql"
	"go/go-server-boilerplate/models"
	"log"
)

type userRepository struct {
	db *sql.DB
}

// UserRepository ...
type UserRepository interface {
	GetAllUsers() []models.UserDTO
	CreateUser(user *models.User) int
	GetUser(id int) (models.UserDTO, error)
	UpdateUser(id int, user *models.UserDTO) error
	DeleteUser(id int) error
}

// NewUserRepository ...
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetAllUsers : get a list of users
func (u *userRepository) GetAllUsers() []models.UserDTO {
	rows, err := u.db.Query("SELECT name, email, phone from users")

	if err != nil {
		log.Fatal("Failed with", err)
	}
	defer rows.Close()

	users := []models.UserDTO{}

	for rows.Next() {
		u := models.UserDTO{}
		if err := rows.Scan(&u.Name, &u.Phone, &u.Email); err != nil {
			log.Fatal("Failed with", err)
		}
		users = append(users, u)
	}
	return users
}

// CreateUser : create a user
func (u *userRepository) CreateUser(user *models.User) int {
	err := u.db.QueryRow(
		"INSERT INTO users(name, username, phone, email) VALUES($1, $2, $3, $4) RETURNING id",
		user.Name, user.Username, user.Phone, user.Email).Scan(&user.ID)
	if err != nil {
		log.Fatal(err)
	}
	return user.ID
}

// GetUser : get a user by id
func (u *userRepository) GetUser(id int) (models.UserDTO, error) {
	var userDTO models.UserDTO
	err := u.db.QueryRow(
		"SELECT  name, email, phone FROM users where id=$1", id).Scan(&userDTO.Name, &userDTO.Email, &userDTO.Phone)
	return userDTO, err
}

// UpdateUser : update a user by id
func (u *userRepository) UpdateUser(id int, user *models.UserDTO) error {
	_, err := u.db.Exec(
		"UPDATE users SET name=$1, email=$2, phone=$3 where id=$4", user.Name, user.Email, user.Phone, id)
	return err
}

// DeleteUser : delete a user by id
func (u *userRepository) DeleteUser(id int) error {
	_, err := u.db.Exec("DELETE from users where id=$1", id)
	return err
}
