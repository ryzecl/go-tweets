package comment

import (
	"context"
	"go-tweets/internal/model"
)

func (r *commentRepository) StoreLikeComment(ctx context.Context, model *model.CommentLikeModel) error {
	query := `INSERT INTO comment_likes (comment_id, user_id, created_at) VALUES (?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.CommentID, model.UserID, model.CreatedAt)
	return err
}
