package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/common/model"
)

// Post represents the client for posts table
type Post struct{}

// Create creates and returns single user
func (p Post) Create(db *sql.DB, post model.Post) (model.Post, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	statement, err := db.PrepareContext(ctx, `
		INSERT INTO posts
		(author_id, title, body, likes, created_at, updated_at, deleted_at)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`,
	)
	if err != nil {
		return model.Post{}, err
	}

	post.Base.BeforeCreate()

	result, err := statement.ExecContext(ctx, post.AuthorID, post.Title, post.Body, post.Likes, post.CreatedAt, post.UpdatedAt, post.DeletedAt)
	if err != nil {
		return model.Post{}, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return model.Post{}, err
	}

	return p.FindByID(db, insertedID)
}

// FindByID returns single post by ID
func (p Post) FindByID(db *sql.DB, id int64) (model.Post, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var post model.Post
	row := db.QueryRowContext(ctx,
		`SELECT * FROM posts WHERE id = (?)`, id,
	)

	if err := row.Scan(
		&post.ID,
		&post.AuthorID,
		&post.Title,
		&post.Body,
		&post.Likes,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.DeletedAt,
	); err != nil {
		fmt.Println("errrrr", err)
		return model.Post{}, nil
	}

	return post, nil
}
