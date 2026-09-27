package post

import (
	"context"
	"errors"
	"go-tweets/internal/model"
)

func (r *postRepository) UpdatePost(ctx context.Context, model *model.PostModel, postID int64) error {
	query := `UPDATE posts SET title = ?,  content = ?, updated_at = ? WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, model.Title, model.Content, model.UpdatedAt, postID)
	if err != nil {
		return err
	}

	rowAffected, _ := result.RowsAffected()
	if rowAffected == 0 {
		return errors.New("Nothing to update")
	}

	return nil
}
