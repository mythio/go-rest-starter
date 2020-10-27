package user

// Service is an interface for userService
type Service interface {
	Signup(*User) error
	Signin(*User) error
}

type 