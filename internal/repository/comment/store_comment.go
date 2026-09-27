package comment

import (
	"context"
	"go-tweets/internal/model"
)

func (r *commentRepository) StoreComment(ctx context.Context, model *model.CommentModel) error {
	query := `INSERT INTO comments (post_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.Content, model.CreatedAt, model.UpdatedAt)
	return err
}
