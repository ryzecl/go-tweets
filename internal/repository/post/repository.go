package post

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
	"time"
)

type PostRepository interface {
	StorePost(ctx context.Context, model *model.PostModel) (int64, error)
	GetPostByID(ctx context.Context, postID int64) (*model.PostWithUserModel, error)
	UpdatePost(ctx context.Context, model *model.PostModel, postID int64) error
	SoftDeletePost(ctx context.Context, postID int64, now time.Time) error
	IsUserAlreadyLikePost(ctx context.Context, postID, userID int64) (bool, error)
	DeleteLikePost(ctx context.Context, postID, userID int64) error
	StoreLikePost(ctx context.Context, model *model.PostLikeModel) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}
