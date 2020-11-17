package post

import (
	"database/sql"

	"github.com/mythio/go-rest-starter/pkg/api/routes/post/platform/mysql"
	"github.com/mythio/go-rest-starter/pkg/common/model"
	"github.com/mythio/go-rest-starter/pkg/common/util/pagination"
)

// Post represents post application service
type Post struct {
	db            *sql.DB
	postRepo      RepositoryPost
	postToTagRepo RepositoryPostToTag
}

// Service represents post application interface
type Service interface {
	Create(ReqCreate) (ResCreate, error)
	Get(int64) (ResGet, error)
	GetAll(pagination.ReqPagination) (ResGetAll, error)
	Update(ReqUpdate) (ResUpdate, error)
}

// RepositoryPost represents post repository interface
type RepositoryPost interface {
	Create(*sql.DB, model.Post) (model.Post, error)
	FindByID(*sql.DB, int64) (model.Post, error)
	GetAll(db *sql.DB, page pagination.Pagination) ([]model.Post, error)
	Update(*sql.DB, model.Post) (model.Post, error)
}

// RepositoryPostToTag represents post repository interface
type RepositoryPostToTag interface {
	Create(db *sql.DB, postToTag model.PostToTag) error
	GetTagIDs(db *sql.DB, id int64) ([]int64, error)
}

func newPostApplicationService(db *sql.DB, repo RepositoryPost) *Post {
	return &Post{
		db:       db,
		postRepo: repo,
	}
}

// InitService initializes auth application service
func InitService(db *sql.DB) *Post {
	return newPostApplicationService(db, mysql.Post{})
}
