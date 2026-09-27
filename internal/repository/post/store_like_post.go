package post

import (
	"context"
	"go-tweets/internal/model"
)

func (r *postRepository) StoreLikePost(ctx context.Context, model *model.PostLikeModel) error {
	query := `INSERT INTO post_likes (post_id, user_id, created_at) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CreatedAt)

	return err
}
