package post

import (
	"context"
	"errors"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *postService) LikeOrUnlikePost(ctx context.Context, postID, userID int64) (int, error) {
	// check post was exists
	postExists, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if postExists == nil {
		return http.StatusNotFound, errors.New("Tweets not found")
	}

	// check user already like or not
	isUserAlreadyLikePost, err := s.postRepo.IsUserAlreadyLikePost(ctx, postID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// if user already like, delete data
	if isUserAlreadyLikePost {
		err := s.postRepo.DeleteLikePost(ctx, postID, userID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		now := time.Now()
		err := s.postRepo.StoreLikePost(ctx, &model.PostLikeModel{
			PostID:    postID,
			UserID:    userID,
			CreatedAt: now,
		})
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}
