package model

// User represents user domain model
type User struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	Base
}
