package post

import (
	"context"
	"go-tweets/internal/config"
	"go-tweets/internal/dto"
	"go-tweets/internal/repository/post"
)

type PostService interface {
	CreatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID int64) (int64, int, error)
	UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, postID, userID int64) (int, error)
}

type postService struct {
	cfg      *config.Config
	postRepo post.PostRepository
}

func NewPostService(cfg *config.Config, postRepo post.PostRepository) PostService {
	return &postService{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
