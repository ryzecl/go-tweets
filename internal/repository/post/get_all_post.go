package post

import (
	"context"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
)

func (r *postRepository) GetAllPost(ctx context.Context, param *dto.GetAllPostRequest, offset int) ([]model.PostWithUserModel, error) {
	query := `SELECT p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, u.username, COUNT(pl.id) as like_count
	FROM posts as p
	JOIN users as u ON u.id = p.user_id
	LEFT JOIN post_likes as pl ON pl.post_id = p.id
	WHERE p.deleted_at IS NULL
	GROUP BY p.id, p.title, p.content, p.user_id, p.created_at, p.updated_at, u.username
	ORDER BY p.created_at DESC
	LIMIT ?
	OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, param.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.PostWithUserModel, 0)
	for rows.Next() {
		var data model.PostWithUserModel
		err := rows.Scan(&data.ID, &data.Title, &data.Content, &data.UserID, &data.CreatedAt, &data.UpdatedAt, &data.Username, &data.LikeCount)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}
