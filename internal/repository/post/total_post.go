package post

import "context"

func (r *postRepository) TotalPost(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(id) FROM posts
	WHERE deleted_at IS NULL`

	var total int64
	err := r.db.QueryRowContext(ctx, query).Scan(&total)

	return total, err
}
