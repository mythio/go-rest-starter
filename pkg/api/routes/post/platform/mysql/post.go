package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/common/model"
	"github.com/mythio/go-rest-starter/pkg/common/util/pagination"
)

// Post represents the repository for posts table
type Post struct{}

// Create creates and returns single post
func (p Post) Create(db *sql.DB, post model.Post) (model.Post, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	statement, err := db.PrepareContext(ctx, `
		INSERT INTO posts
		(author_id, title, body, likes, created_at, updated_at, deleted_at)
		VALUES(?, ?, ?, ?, ?, ?, ?)
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

	post := model.Post{}
	row := db.QueryRowContext(ctx, `
		SELECT *
		FROM posts 
		WHERE id = (?)
	`, id,
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

// GetAll returns a paginated list of the posts
func (p Post) GetAll(db *sql.DB, page pagination.Pagination) ([]model.Post, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	posts := []model.Post{}
	rows, err := db.QueryContext(ctx, `
		SELECT *
		FROM posts
		WHERE deleted_at = 0
		LIMIT ? OFFSET ?
	`, page.Limit, page.Offset,
	)
	if err != nil {
		return []model.Post{}, err
	}

	for rows.Next() {
		post := model.Post{}
		if err := rows.Scan(
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
			return []model.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update updates and returns single post by ID
func (p Post) Update(db *sql.DB, post model.Post) (model.Post, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	statement, err := db.PrepareContext(ctx, `
		UPDATE posts
		SET title = (?), body = (?), likes = (?), updated_at = (?)
		WHERE id = (?)
		`,
	)
	if err != nil {
		return model.Post{}, err
	}

	post.Base.BeforeUpdate()

	result, err := statement.ExecContext(ctx, post.Title, post.Body, post.Likes, post.UpdatedAt, post.ID)
	if err != nil {
		return model.Post{}, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return model.Post{}, err
	}

	return p.FindByID(db, insertedID)
}
