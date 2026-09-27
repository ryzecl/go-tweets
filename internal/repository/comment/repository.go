package comment

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

type CommentRepository interface {
	StoreComment(ctx context.Context, model *model.CommentModel) error
	DetailComment(ctx context.Context, commentID int64) (*model.CommentModel, error)
	IsUserAlreadyLikeComment(ctx context.Context, commentID, userID int64) (bool, error)
	DeleteLikeComment(ctx context.Context, commentID, userID int64) error
	StoreLikeComment(ctx context.Context, model *model.CommentLikeModel) error
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}
