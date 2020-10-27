package user

import (
	"fmt"

	"github.com/mythio/go-rest-starter/util/logger"
	"gopkg.in/go-playground/validator.v9"
)

type userService struct {
	userRepo  Repository
	logger    logger.Logger
	validator *validator.Validate
}

// NewUserService creates and returns a new User service
func NewUserService(userRepo Repository, logger logger.Logger, validate *validator.Validate) UserService {
	return &userService{
		userRepo,
		logger,
		validate,
	}
}

func (s *userService) Signup(user *User) error {
	if err := user.validate(s.validator); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		return err
	}

	return s.userRepo.Create(user)
}

func (s *userService) Signin(user *User) error {
	if err := user.validate(s.validator); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		return err
	}

	return s.userRepo.Create(user)
}
