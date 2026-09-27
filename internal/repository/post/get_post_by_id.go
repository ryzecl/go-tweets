package post

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

func (r *postRepository) GetPostByID(ctx context.Context, postID int64) (*model.PostModel, error) {
	query := `SELECT id, title, content, user_id, created_at, updated_at FROM posts WHERE id = ? AND deleted_at IS NULL`

	row := r.db.QueryRowContext(ctx, query, postID)
	var result model.PostModel
	err := row.Scan(&result.ID, &result.Title, &result.Content, &result.UserID, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
