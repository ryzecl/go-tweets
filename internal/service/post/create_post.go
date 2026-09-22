package post

import (
	"context"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *postService) CreatePost(ctx context.Context, req *dto.CreatePostRequest, userID int64) (int64, int, error) {
	// Store postRepo
	now := time.Now()
	insertedID, err := s.postRepo.StorePost(ctx, &model.PostModel{
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	// Return
	return insertedID, http.StatusCreated, nil
}
