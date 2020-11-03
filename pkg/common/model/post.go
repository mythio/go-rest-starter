package model

// Post represents post domain model
type Post struct {
	AuthorID int64  `json:"author_id" db:"author_id"`
	Title    string `json:"title" db:"title"`
	Body     string `json:"body" db:"body"`
	Likes    int64  `json:"likes" db:"likes"`
	Base
}
