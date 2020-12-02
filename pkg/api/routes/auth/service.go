package auth

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
	"github.com/mythio/go-rest-starter/pkg/api/routes/auth/platform/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/model"
)

// Auth represents user application service
type Auth struct {
	db   *sql.DB
	repo Repository
	sec  Security
	tk   Token
}

// Service represents user application interface
type Service interface {
	Signup(ReqSignup) (ResSignup, error)
	Signin(ReqSignin) (ResSignin, error)
	Me(int64) (ResMe, error)
}

// Repository represents user repository interface
type Repository interface {
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

// New creates new user application service
func newUserApplicationService(db *sql.DB, repo Repository, sec Security, tk Token) *Auth {
	return &Auth{
		db:   db,
		repo: repo,
		sec:  sec,
		tk:   tk,
	}
}

// InitService initializes auth application service
func InitService(db *sql.DB, sec Security, tk Token) *Auth {
	return newUserApplicationService(db, mysql.User{}, sec, tk)
}
