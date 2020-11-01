package auth

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
	"github.com/mythio/go-rest-starter/pkg/api/auth/platform/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/model"
	"github.com/mythio/go-rest-starter/pkg/common/model/res"
	"github.com/mythio/go-rest-starter/pkg/common/util/logger"
)

// Service represents user application interface
type Service interface {
	Signup(model.User) (res.AuthUser, error)
	Signin(model.User) (res.AuthUser, error)
	Me(int64) (model.User, error)
}

// UserRepo represents user repository interface
type UserRepo interface {
	Create(*sql.DB, model.User) (model.User, error)
	FindByID(*sql.DB, int64) (model.User, error)
	FindByEmail(*sql.DB, string) (model.User, error)
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
	db    *sql.DB
	uRepo UserRepo
	sec   Security
	log   logger.Logger
	tk    Token
}

// New creates new user application service
func newUserApplicationService(db *sql.DB, repo UserRepo, sec Security, log logger.Logger, tk Token) *Auth {
	return &Auth{
		db:    db,
		uRepo: repo,
		sec:   sec,
		log:   log,
		tk:    tk,
	}
}

// InitService initializes auth application service
func InitService(db *sql.DB, sec Security, log logger.Logger, tk Token) *Auth {
	return newUserApplicationService(db, mysql.User{}, sec, log, tk)
}
