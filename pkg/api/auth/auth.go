package auth

import (
	"errors"

	"github.com/mythio/go-rest-starter/pkg/common/model"
	"gorm.io/gorm"
)

// Signup creates a new user account
func (u *Auth) Signup(user model.User) (model.User, error) {
	existingUser, err := u.uRepo.FindByEmail(u.db, user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return model.User{}, err
	}

	if existingUser != (model.User{}) {
		return model.User{}, errors.New("Email id used")
	}

	user.Password = u.sec.Hash(user.Password)
	return u.uRepo.Create(u.db, user)
}

// Signin gets existing user account
func (u *Auth) Signin(user model.User) (model.User, error) {
	existingUser, err := u.uRepo.FindByEmail(u.db, user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return model.User{}, err
	}

	if existingUser == (model.User{}) {
		return model.User{}, errors.New("not found")
	}

	res := u.sec.Verify(existingUser.Password, user.Password)

	if res == false {
		return model.User{}, errors.New("incorrect password")
	}

	return existingUser, nil
}

// View returns single user
func (u *Auth) View(id uint32) (model.User, error) {
	return u.uRepo.FindByID(u.db, id)
}

// Update updates user's contact information
func (u *Auth) Update(user model.User) (model.User, error) {
	u.uRepo.Update(u.db, user)

	return u.uRepo.FindByID(u.db, user.ID)
}

// Delete deletes a user
func (u *Auth) Delete(id uint32) error {
	return u.uRepo.Delete(u.db, id)
}
