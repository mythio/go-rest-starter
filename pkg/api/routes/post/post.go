package post

import (
	"github.com/mythio/go-rest-starter/pkg/common/model"
	"github.com/mythio/go-rest-starter/pkg/common/util/pagination"
)

type postAttr struct {
	ID        int64   `json:"id"`
	AuthorID  int64   `json:"author_id"`
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	Likes     int64   `json:"likes"`
	Tags      []int64 `json:"tags"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

// ReqCreate defines the request recieved by create service
type ReqCreate struct {
	AuthorID int64
	Title    string
	Body     string
	Tags     []int64
}

// ResCreate defines the response sent by create service
type ResCreate postAttr

// ReqUpdate defines the request recieved by update service
type ReqUpdate struct {
	ReqCreate
	ID int64
}

// ResUpdate defines the response sent by update service
type ResUpdate postAttr

// ResGet defines the response sent by get service
type ResGet postAttr

// ResGetAll defines the response sent by the getAll service
type ResGetAll []postAttr

// Create creates a new post for the user
func (p *Post) Create(post ReqCreate) (ResCreate, error) {
	_post, err := p.postRepo.Create(p.db, model.Post{
		AuthorID: post.AuthorID,
		Title:    post.Title,
		Body:     post.Body,
		Likes:    0,
	})

	for _, tag := range post.Tags {
		if err := p.postToTagRepo.Create(p.db, model.PostToTag{
			PostID: _post.ID,
			TagID:  tag,
		}); err != nil {
			return ResCreate{}, err
		}
	}

	_tags, err := p.postToTagRepo.GetTagIDs(p.db, _post.ID)
	if err != nil {
		return ResCreate{}, err
	}

	if err != nil {
		return ResCreate{}, err
	}

	result := ResCreate{}

	result.ID = _post.ID
	result.AuthorID = _post.AuthorID
	result.Title = _post.Title
	result.Body = _post.Body
	result.Likes = _post.Likes
	result.Tags = _tags
	result.CreatedAt = _post.CreatedAt
	result.UpdatedAt = _post.UpdatedAt

	return result, err
}

// Update updates a post with the given id
func (p *Post) Update(post ReqUpdate) (ResUpdate, error) {
	_, err := p.postRepo.FindByID(p.db, post.ID)
	if err != nil {
		return ResUpdate{}, err
	}

	postModel := model.Post{}

	postModel.ID = post.ID
	postModel.AuthorID = post.AuthorID
	postModel.Title = post.Title
	postModel.Body = post.Body
	postModel.Likes = 0

	_post, err := p.postRepo.Update(p.db, postModel)
	if err != nil {
		return ResUpdate{}, err
	}

	result := ResUpdate{}

	result.ID = _post.ID
	result.AuthorID = _post.AuthorID
	result.Title = _post.Title
	result.Body = _post.Body
	result.Likes = _post.Likes
	result.CreatedAt = _post.CreatedAt
	result.UpdatedAt = _post.UpdatedAt

	return result, nil
}

// Get returns a post with the given id
func (p *Post) Get(id int64) (ResGet, error) {
	_post, err := p.postRepo.FindByID(p.db, id)
	if err != nil {
		return ResGet{}, err
	}

	result := ResGet{}

	result.ID = _post.ID
	result.AuthorID = _post.AuthorID
	result.Title = _post.Title
	result.Body = _post.Body
	result.Likes = _post.Likes
	result.CreatedAt = _post.CreatedAt
	result.UpdatedAt = _post.UpdatedAt

	return result, nil
}

// GetAll returns a paginated list of posts
func (p *Post) GetAll(page pagination.ReqPagination) (ResGetAll, error) {
	pageDb := page.Transform()
	_posts, err := p.postRepo.GetAll(p.db, pageDb)
	if err != nil {
		return ResGetAll{}, err
	}

	result := ResGetAll{}

	for _, _post := range _posts {
		post := postAttr{}

		post.ID = _post.ID
		post.AuthorID = _post.AuthorID
		post.Title = _post.Title
		post.Body = _post.Body
		post.Likes = _post.Likes
		post.CreatedAt = _post.CreatedAt
		post.UpdatedAt = _post.UpdatedAt

		result = append(result, post)
	}

	return result, nil
}
