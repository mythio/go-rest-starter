package model

// Post represents post domain model
type Post struct {
	AuthorID int64
	Title    string
	Body     string
	Likes    int64
	Base
}
