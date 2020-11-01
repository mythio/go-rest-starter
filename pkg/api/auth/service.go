package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mythio/go-rest-starter/pkg/api/auth/platform/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/model"
	"github.com/mythio/go-rest-starter/pkg/common/model/res"
	"github.com/mythio/go-rest-starter/pkg/common/util/logger"
	"gorm.io/gorm"
)

// Service represents user application interface
type Service interface {
	Signup(model.User) (res.AuthUser, error)
	Signin(model.User) (res.AuthUser, error)
	Me(uint32) (model.User, error)
}

// UserRepo represents user repository interface
type UserRepo interface {
	Create(*gorm.DB, model.User) (model.User, error)
	FindByID(*gorm.DB, uint32) (model.User, error)
	FindByEmail(*gorm.DB, string) (model.User, error)
}

// Security represents security interface
type Security interface {
	Hash(password string) string
	Verify(hashedPassword string, password string) bool
}

// Token represents jwt interface
type Token interface {
	GenerateToken(model.User) (string, error)
	ParseToken(string) (*jwt.Token, error)
}

// Auth represents user application service
type Auth struct {
	db    *gorm.DB
	uRepo UserRepo
	sec   Security
	log   logger.Logger
	tk    Token
}

// New creates new user application service
func newUserApplicationService(db *gorm.DB, repo UserRepo, sec Security, log logger.Logger, tk Token) *Auth {
	return &Auth{
		db:    db,
		uRepo: repo,
		sec:   sec,
		log:   log,
		tk:    tk,
	}
}

// InitService initializes auth application service
func InitService(db *gorm.DB, sec Security, log logger.Logger, tk Token) *Auth {
	return newUserApplicationService(db, mysql.User{}, sec, log, tk)
}
