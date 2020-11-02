package auth

import (
	"errors"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/common/model"
	"github.com/mythio/go-rest-starter/pkg/common/model/res"
)

// Signup creates a new user account
func (a *Auth) Signup(user model.User) (res.AuthUser, error) {
	existingUser, err := a.uRepo.FindByEmail(a.db, user.Email)
	if err != nil {
		return res.AuthUser{}, err
	}

	if existingUser != (model.User{}) {
		return res.AuthUser{}, errors.New("Email id used")
	}

	user.Password = a.sec.Hash(user.Password)
	createdUser, err := a.uRepo.Create(a.db, user)
	if err != nil {
		return res.AuthUser{}, err
	}

	accessToken, err := a.tk.GenerateToken(createdUser)
	if err != nil {
		return res.AuthUser{}, err
	}

	return (res.AuthUser{
		User:        createdUser,
		AccessToken: accessToken,
	}), nil
}

// Signin gets existing user account
func (a *Auth) Signin(user model.User) (res.AuthUser, error) {
	existingUser, err := a.uRepo.FindByEmail(a.db, user.Email)
	if err != nil {
		fmt.Println("errrrr", err)
		return res.AuthUser{}, err
	}

	if existingUser == (model.User{}) {
		return res.AuthUser{}, errors.New("not found")
	}

	if isValid := a.sec.Verify(existingUser.Password, user.Password); !isValid {
		return res.AuthUser{}, errors.New("incorrect password")
	}

	accessToken, err := a.tk.GenerateToken(existingUser)
	if err != nil {
		return res.AuthUser{}, err
	}

	return (res.AuthUser{
		User:        existingUser,
		AccessToken: accessToken,
	}), nil
}

// Me returns the current user
func (a *Auth) Me(id int64) (model.User, error) {
	return a.uRepo.FindByID(a.db, id)
}
