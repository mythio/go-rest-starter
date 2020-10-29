package auth

import (
	"github.com/mythio/go-rest-starter/pkg/api/auth/platform/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/model"
	"github.com/mythio/go-rest-starter/pkg/common/util/logger"
	"gorm.io/gorm"
)

// Service represents user application interface
type Service interface {
	Signup(model.User) (model.User, error)
	Signin(model.User) (model.User, error)
	View(uint32) (model.User, error)
	Update(model.User) (model.User, error)
	Delete(uint32) error
}

// UserRepo represents user repository interface
type UserRepo interface {
	Create(*gorm.DB, model.User) (model.User, error)
	FindByID(*gorm.DB, uint32) (model.User, error)
	FindByEmail(*gorm.DB, string) (model.User, error)
	Update(*gorm.DB, model.User) error
	Delete(*gorm.DB, uint32) error
}

// Security represents security interface
type Security interface {
	Hash(password string) string
	Verify(hashedPassword string, password string) bool
}

// Auth represents user application service
type Auth struct {
	db    *gorm.DB
	uRepo UserRepo
	sec   Security
	log   logger.Logger
}

// New creates new user application service
func newUserApplicationService(db *gorm.DB, repo UserRepo, sec Security, log logger.Logger) *Auth {
	return &Auth{
		db:    db,
		uRepo: repo,
		sec:   sec,
		log:   log,
	}
}

func InitService(db *gorm.DB, sec Security, log logger.Logger) *Auth {
	return newUserApplicationService(db, mysql.User{}, sec, log)
}
