package comment

import "context"

func (r *commentRepository) DeleteLikeComment(ctx context.Context, commentID, userID int64) error {
	query := `DELETE FROM comment_likes WHERE comment_id = ? AND user_id = ?`
	_, err := r.db.ExecContext(ctx, query, commentID, userID)
	return err
}
