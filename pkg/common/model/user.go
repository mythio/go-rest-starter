package model

// User represents user domain model
type User struct {
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	Base
}
