package mysql

import (
	"context"
	"database/sql"

	"github.com/mythio/go-rest-starter/pkg/common/model"
)

// PostsToTags represents the repository for posts_to_tags table
type PostsToTags struct{}

// Create creates a single post to tag
func (p PostsToTags) Create(db *sql.DB, postToTag model.PostToTag) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	statement, err := db.PrepareContext(ctx, `
		INSERT INTO posts_to_tags
		(post_id, tag_id)
		VALUES(?, ?)
	`,
	)
	if err != nil {
		return err
	}

	result, err := statement.ExecContext(ctx, postToTag.PostID, postToTag.TagID)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows != 1 {
		return err
	}

	return nil
}

// GetTagIDs returns all tags for a post
func (p PostsToTags) GetTagIDs(db *sql.DB, id int64) ([]int64, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT tag_id
		FROM posts_to_tags
		WHERE post_id = (?)
	`, id)
	if err != nil {
		return []int64{}, err
	}

	var tagIDs []int64

	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return []int64{}, err
		}

		tagIDs = append(tagIDs, id)
	}

	return tagIDs, nil
}
