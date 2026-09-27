package post

import (
	"context"
	"database/sql"
)

func (r *postRepository) IsUserAlreadyLikePost(ctx context.Context, postID, userID int64) (bool, error) {
	query := `SELECT id FROM post_likes WHERE post_id = ? AND user_id = ?`
	row := r.db.QueryRowContext(ctx, query, postID, userID)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
