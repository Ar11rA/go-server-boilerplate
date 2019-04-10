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
