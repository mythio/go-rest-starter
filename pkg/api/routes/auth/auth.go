package auth

import (
	"errors"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/common/model"
)

type userAttr struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// ResMe defines the user sent by service
type ResMe userAttr

// ReqSignup defines the user recieved by service
type ReqSignup struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

// ResSignup defines the authenticated user sent by service
type ResSignup ResSignin

// ReqSignin defines the user received by service
type ReqSignin struct {
	Email    string
	Password string
}

// ResSignin defines the authenticated user sent by service
type ResSignin struct {
	User        userAttr
	AccessToken string `json:"access_token"`
}

// Signup creates a new user account
func (a *Auth) Signup(user ReqSignup) (ResSignup, error) {
	_user, err := a.repo.FindByEmail(a.db, user.Email)
	if err != nil {
		return ResSignup{}, err
	}

	if _user != (model.User{}) {
		return ResSignup{}, errors.New("Email id used")
	}

	u := model.User{}
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Username = user.Username
	u.Email = user.Email
	u.Password = a.sec.Hash(user.Password)

	_user, err = a.repo.Create(a.db, u)
	if err != nil {
		return ResSignup{}, err
	}

	accessToken, err := a.tk.GenerateToken(_user)
	if err != nil {
		return ResSignup{}, err
	}

	result := ResSignup{}

	result.User.ID = _user.ID
	result.User.FirstName = _user.FirstName
	result.User.LastName = _user.LastName
	result.User.Username = _user.Username
	result.User.Email = _user.Email
	result.User.CreatedAt = _user.CreatedAt
	result.User.UpdatedAt = _user.UpdatedAt

	result.AccessToken = accessToken

	return result, nil
}

// Signin gets existing user account
func (a *Auth) Signin(user ReqSignin) (ResSignin, error) {
	_user, err := a.repo.FindByEmail(a.db, user.Email)
	if err != nil {
		fmt.Println("errrrr", err)
		return ResSignin{}, err
	}

	if _user == (model.User{}) {
		return ResSignin{}, errors.New("not found")
	}

	if isValid := a.sec.Verify(_user.Password, user.Password); !isValid {
		return ResSignin{}, errors.New("incorrect password")
	}

	accessToken, err := a.tk.GenerateToken(_user)
	if err != nil {
		return ResSignin{}, err
	}

	result := ResSignin{}

	result.User.ID = _user.ID
	result.User.FirstName = _user.FirstName
	result.User.LastName = _user.LastName
	result.User.Username = _user.Username
	result.User.Email = _user.Email
	result.User.CreatedAt = _user.CreatedAt
	result.User.UpdatedAt = _user.UpdatedAt

	result.AccessToken = accessToken

	return result, nil
}

// Me returns the current user
func (a *Auth) Me(id int64) (ResMe, error) {
	_user, err := a.repo.FindByID(a.db, id)
	if err != nil {
		return ResMe{}, err
	}

	result := ResMe{}

	result.ID = _user.ID
	result.FirstName = _user.FirstName
	result.LastName = _user.LastName
	result.Username = _user.Username
	result.Email = _user.Email
	result.CreatedAt = _user.CreatedAt
	result.UpdatedAt = _user.UpdatedAt

	return result, nil
}
