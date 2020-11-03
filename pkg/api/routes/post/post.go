package post

import "github.com/mythio/go-rest-starter/pkg/common/model"

// ReqCreate defines the post sent by service
type ReqCreate struct {
	AuthorID int64
	Title    string
	Body     string
	Tags     []int64
}

// ResCreate defines the post sent by service
type ResCreate struct {
	ID        int64
	AuthorID  int64
	Title     string
	Body      string
	Likes     int64
	Tags      []int64
	CreatedAt int64
	UpdatedAt int64
}

// ReqUpdate defines the post sent by service
type ReqUpdate ReqCreate

// ResUpdate defines the post sent by service
type ResUpdate ResCreate

// Create creates a new post for the user
func (p *Post) Create(post ReqCreate) (ResCreate, error) {
	_post, err := p.repo.Create(p.db, model.Post{
		AuthorID: post.AuthorID,
		Title:    post.Title,
		Body:     post.Body,
		Likes:    0,
	})
	if err != nil {
		return ResCreate{}, err
	}

	result := ResCreate{}

	result.ID = _post.ID
	result.AuthorID = _post.AuthorID
	result.Title = _post.Title
	result.Body = _post.Body
	result.Likes = _post.Likes
	result.CreatedAt = _post.CreatedAt
	result.UpdatedAt = _post.UpdatedAt

	return result, err
}
