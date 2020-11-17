package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mythio/go-rest-starter/pkg/common/model"
)

// Tag represents the repository for tags table
type Tag struct{}

// Create creates and returns single tag
func (t Tag) Create(db *sql.DB, tag model.Tag) (model.Tag, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	statement, err := db.PrepareContext(ctx, `
		INSERT INTO tags
		(name, description, created_at, updated_at, deleted_at)
		VALUES(?, ?, ?, ?, ?)
		`,
	)
	if err != nil {
		return model.Tag{}, err
	}

	tag.Base.BeforeCreate()

	result, err := statement.ExecContext(ctx, tag.Name, tag.Description, tag.CreatedAt, tag.UpdatedAt, tag.DeletedAt)
	if err != nil {
		return model.Tag{}, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return model.Tag{}, err
	}

	return t.FindByID(db, insertedID)
}

// FindByID returns single post by ID
func (t Tag) FindByID(db *sql.DB, id int64) (model.Tag, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tag model.Tag
	row := db.QueryRowContext(ctx,
		`SELECT * FROM tags WHERE id = (?)`, id,
	)

	if err := row.Scan(
		&tag.ID,
		&tag.Name,
		&tag.Description,
		&tag.CreatedAt,
		&tag.UpdatedAt,
		&tag.DeletedAt,
	); err != nil {
		fmt.Println("errrrr", err)
		return model.Tag{}, nil
	}

	return tag, nil
}

// Search returns a list of tags matching the name
func (t Tag) Search(db *sql.DB, str string) ([]model.Tag, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT *
		FROM tags
	`)
	if err != nil {
		return []model.Tag{}, err
	}

	var tags []model.Tag

	for rows.Next() {
		var tag model.Tag
		rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.Description,
			&tag.CreatedAt,
			&tag.UpdatedAt,
			&tag.DeletedAt,
		)

		tags = append(tags, tag)
	}

	return tags, nil
}
