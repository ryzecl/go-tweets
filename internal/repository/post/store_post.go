package post

import (
	"context"
	"go-tweets/internal/model"
)

func (r *postRepository) StorePost(ctx context.Context, model *model.PostModel) (int64, error) {
	query := `INSERT INTO posts (user_id, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, model.UserID, model.Title, model.Content, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}
