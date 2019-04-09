package models

// User model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
