package post

import (
	"database/sql"

	"github.com/mythio/go-rest-starter/pkg/common/model"
)

// Post represents post application service
type Post struct {
	db   *sql.DB
	repo Repository
}

// Service represents post application interface
type Service interface {
	Create(ReqCreate) (ResCreate, error)
	Delete(int64) error
}

// Repository represents post repository interface
type Repository interface {
	Create(*sql.DB, model.Post) (model.Post, error)
	FindById(int64) (*sql.DB, model.Post, error)
}

// InitService initializes auth application service
func InitService(db *sql.DB, repo Repository) *Post {
	return &Post{
		db:   db,
		repo: repo,
	}
}
