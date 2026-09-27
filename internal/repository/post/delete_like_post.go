package post

import "context"

func (r *postRepository) DeleteLikePost(ctx context.Context, postID, userID int64) error {
	query := `DELETE FROM post_likes WHERE post_id = ? AND user_id = ?`
	_, err := r.db.ExecContext(ctx, query, postID, userID)
	return err
}
