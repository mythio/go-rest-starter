package user

import (
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	EmailID   string `json:"email_id" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	UUID      string `json:"uuid"`
	ID        uint32 `json:"id"`
	CreatedAt int64  `json:"created_at"`
}

func (u *User) validate(validate *validator.Validate) error {
	err := validate.Struct(u)
	if err != nil {
		return err
	}

	return nil
}
