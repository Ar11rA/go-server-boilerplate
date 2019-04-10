package repository

import (
	"go/go-server-boilerplate/config"
	"go/go-server-boilerplate/models"
	"log"
)

// GetAllUsers : get a list of users
func GetAllUsers() []models.UserDTO {
	rows, err := config.DB.Query("SELECT name, email, phone from users")

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
func CreateUser(user *models.User) int {
	err := config.DB.QueryRow(
		"INSERT INTO users(name, username, phone, email) VALUES($1, $2, $3, $4) RETURNING id",
		user.Name, user.Username, user.Phone, user.Email).Scan(&user.ID)
	if err != nil {
		log.Fatal(err)
	}
	return user.ID
}

// GetUser : get a user by id
func GetUser(id int) (models.UserDTO, error) {
	var userDTO models.UserDTO
	err := config.DB.QueryRow(
		"SELECT  name, email, phone FROM users where id=$1", id).Scan(&userDTO.Name, &userDTO.Email, &userDTO.Phone)
	return userDTO, err
}

// UpdateUser : update a user by id
func UpdateUser(id int, user *models.UserDTO) error {
	_, err := config.DB.Exec(
		"UPDATE users SET name=$1, email=$2, phone=$3 where id=$4", user.Name, user.Email, user.Phone, id)
	return err
}

// DeleteUser : delete a user by id
func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE from users where id=$1", id)
	return err
}
