package post

import (
	"context"
	"go-tweets/internal/config"
	"go-tweets/internal/dto"
	"go-tweets/internal/repository/comment"
	"go-tweets/internal/repository/post"
)

type PostService interface {
	CreatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID int64) (int64, int, error)
	UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, postID, userID int64) (int, error)
	DeletePost(ctx context.Context, postID, userID int64) (int, error)
	LikeOrUnlikePost(ctx context.Context, postID, userID int64) (int, error)
	DetailPost(ctx context.Context, postID int64) (*dto.DetailPostResponse, int, error)
	GetAllPost(ctx context.Context, req *dto.GetAllPostRequest) (*dto.GetAllPostResponse, int, error)
}

type postService struct {
	cfg         *config.Config
	postRepo    post.PostRepository
	commentRepo comment.CommentRepository
}

func NewPostService(cfg *config.Config, postRepo post.PostRepository, commentRepo comment.CommentRepository) PostService {
	return &postService{
		cfg:         cfg,
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}
