package post

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *postService) UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, postID, userID int64) (int, error) {
	// check post was exists
	postExists, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if postExists == nil {
		return http.StatusNotFound, errors.New("Tweets not found")
	}

	if postExists.UserID != userID {
		return http.StatusNotFound, errors.New("Tweets not found")
	}

	// update post
	err = s.postRepo.UpdatePost(ctx, &model.PostModel{
		Title:     req.Title,
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// return
	return http.StatusOK, nil
}
