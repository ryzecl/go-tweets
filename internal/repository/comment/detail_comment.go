package comment

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

func (r *commentRepository) DetailComment(ctx context.Context, commentID int64) (*model.CommentModel, error) {
	query := `SELECT id, post_id, user_id, content, created_at FROM comments WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, commentID)
	var result model.CommentModel
	err := row.Scan(&result.ID, &result.PostID, &result.UserID, &result.Content, &result.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
