package post

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

func (r *postRepository) GetPostByID(ctx context.Context, postID int64) (*model.PostWithUserModel, error) {
	query := `SELECT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, u.username, COUNT(pl.id) as like_count
	FROM posts as p
	JOIN users as u ON p.user_id = u.id
	LEFT JOIN post_likes as pl ON pl.post_id = p.id
	WHERE p.id = ?
	AND p.deleted_at IS NULL
	GROUP BY p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, u.username`

	row := r.db.QueryRowContext(ctx, query, postID)
	var result model.PostWithUserModel
	err := row.Scan(&result.ID, &result.Title, &result.Content, &result.UserID, &result.CreatedAt, &result.UpdatedAt, &result.Username, &result.LikeCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
