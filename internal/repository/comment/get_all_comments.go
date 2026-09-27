package comment

import (
	"context"
	"database/sql"
	"fmt"
	"go-tweets/internal/model"
	"strings"
)

func (r *commentRepository) GetCommentsByPostIDs(ctx context.Context, postIDs []int64) ([]model.CommentModel, error) {
	if len(postIDs) == 0 {
		return []model.CommentModel{}, nil
	}

	placeholders := make([]string, len(postIDs))
	args := make([]interface{}, len(postIDs))
	for i, id := range postIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`SELECT c.id, c.post_id, c.user_id, u.username, c.content, c.created_at, c.updated_at, COUNT(cl.id) as like_count
		FROM comments as c
		JOIN users as u ON u.id = c.user_id
		LEFT JOIN comment_likes as cl ON cl.comment_id = c.id
		WHERE c.post_id IN (%s)
		GROUP BY c.id, c.post_id, c.user_id, u.username, c.content, c.created_at, c.updated_at
		ORDER BY like_count DESC`, strings.Join(placeholders, ","))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return []model.CommentModel{}, nil
		}
		return []model.CommentModel{}, err
	}

	result := make([]model.CommentModel, 0)
	for rows.Next() {
		var data model.CommentModel
		err = rows.Scan(&data.ID, &data.PostID, &data.UserID, &data.Username, &data.Content, &data.CreatedAt, &data.UpdatedAt, &data.LikeCount)
		if err != nil {
			return []model.CommentModel{}, err
		}

		result = append(result, model.CommentModel{
			ID:        data.ID,
			PostID:    data.PostID,
			UserID:    data.UserID,
			Username:  data.Username,
			Content:   data.Content,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
			LikeCount: data.LikeCount,
		})
	}

	return result, nil
}
