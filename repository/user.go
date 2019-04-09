package repository

import (
	"go/go-server-boilerplate/config"
	"go/go-server-boilerplate/models"
	"log"
)

// GetAllUsers : get a list of users
func GetAllUsers() []models.User {
	rows, err := config.DB.Query("SELECT * from users")

	if err != nil {
		log.Fatal("Failed with", err)
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		u := models.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Username, &u.Phone, &u.Email); err != nil {
			log.Fatal("Failed with", err)
		}
		users = append(users, u)
	}
	return users
}
