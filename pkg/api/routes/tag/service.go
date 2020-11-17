package tag

import (
	"database/sql"

	"github.com/mythio/go-rest-starter/pkg/api/routes/tag/platform/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/model"
)

// Tag represents tag service
type Tag struct {
	db   *sql.DB
	repo Repository
}

type Service interface {
	Create(ReqCreate) (ResCreate, error)
	Get(int64) (ResGet, error)
	Search(string) (ResSearch, error)
}

type Repository interface {
	Create(db *sql.DB, tag model.Tag) (model.Tag, error)
	FindByID(db *sql.DB, id int64) (model.Tag, error)
	Search(db *sql.DB, str string) ([]model.Tag, error)
}

func newTagApplicationService(db *sql.DB, repo Repository) *Tag {
	return &Tag{
		db:   db,
		repo: repo,
	}
}

// InitService initializes auth application service
func InitService(db *sql.DB) *Tag {
	return newTagApplicationService(db, mysql.Tag{})
}
