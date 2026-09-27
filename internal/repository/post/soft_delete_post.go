package post

import (
	"context"
	"errors"
	"time"
)

func (r *postRepository) SoftDeletePost(ctx context.Context, postID int64, now time.Time) error {
	query := `UPDATE posts SET deleted_at = ? WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, now, postID)
	if err != nil {
		return err
	}

	rowAffected, _ := result.RowsAffected()
	if rowAffected == 0 {
		return errors.New("Nothing to update")
	}

	return nil
}
