package post

import (
	"context"
	"errors"
	"net/http"
	"time"
)

func (s *postService) DeletePost(ctx context.Context, postID, userID int64) (int, error) {
	// check if the post exists
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

	// soft delete postID
	err = s.postRepo.SoftDeletePost(ctx, postID, time.Now())
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
