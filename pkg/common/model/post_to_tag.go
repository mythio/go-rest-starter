package model

// PostToTag represents posts and tags relationship domain model [pivot table]
type PostToTag struct {
	PostID int64
	TagID  int64
	ID     int64
}
