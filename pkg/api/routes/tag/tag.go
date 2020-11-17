package tag

import "github.com/mythio/go-rest-starter/pkg/common/model"

type tagAttr struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   int64
}

// ReqCreate defines the request recieved by create service
type ReqCreate struct {
	Name        string
	Description string
}

// ResCreate defines the response sent by create service
type ResCreate tagAttr

// ResGet defines the response sent by get service
type ResGet tagAttr

// ResSearch defines the response sent by search service
type ResSearch []tagAttr

// Create creates a new tag
func (t *Tag) Create(tag ReqCreate) (ResCreate, error) {
	_tag, err := t.repo.Create(t.db, model.Tag{
		Name:        tag.Name,
		Description: tag.Description,
	})
	if err != nil {
		return ResCreate{}, nil
	}

	result := ResCreate{}

	result.ID = _tag.ID
	result.Name = _tag.Name
	result.Description = _tag.Description
	result.CreatedAt = _tag.CreatedAt
	result.UpdatedAt = _tag.UpdatedAt
	result.DeletedAt = _tag.DeletedAt

	return result, err
}

// Get returns a tag with the given id
func (t *Tag) Get(id int64) (ResGet, error) {
	_tag, err := t.repo.FindByID(t.db, int64(id))
	if err != nil {
		return ResGet{}, err
	}

	result := ResGet{}

	result.ID = _tag.ID
	result.Name = _tag.Name
	result.Description = _tag.Description
	result.CreatedAt = _tag.CreatedAt
	result.UpdatedAt = _tag.UpdatedAt
	result.DeletedAt = _tag.DeletedAt

	return result, nil
}

// Search returns tags matching the given search string
func (t *Tag) Search(name string) (ResSearch, error) {
	_tags, err := t.repo.Search(t.db, string(name))
	if err != nil {
		return ResSearch{}, err
	}

	result := ResSearch{}

	for _, _tag := range _tags {
		tag := tagAttr{}

		tag.ID = _tag.ID
		tag.Name = _tag.Name
		tag.Description = _tag.Description
		tag.CreatedAt = _tag.CreatedAt
		tag.UpdatedAt = _tag.UpdatedAt
		tag.DeletedAt = _tag.DeletedAt

		result = append(result, tag)
	}

	return result, nil
}
