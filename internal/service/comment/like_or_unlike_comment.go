package comment

import (
	"context"
	"errors"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *commentService) LikeOrUnlikeComment(ctx context.Context, commentID, userID int64) (int, error) {
	// check comment if exists
	commentExists, err := s.commentRepo.DetailComment(ctx, commentID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if commentExists == nil {
		return http.StatusNotFound, errors.New("Comment not found")
	}

	// check user already like commentID
	isUserAlreadyLikeComment, err := s.commentRepo.IsUserAlreadyLikeComment(ctx, commentID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// if user was like, delete LikeOrUnlikeComment
	if isUserAlreadyLikeComment {
		err := s.commentRepo.DeleteLikeComment(ctx, commentID, userID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		now := time.Now()
		err := s.commentRepo.StoreLikeComment(ctx, &model.CommentLikeModel{
			CommentID: commentID,
			UserID:    userID,
			CreatedAt: now,
		})
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	// return
	return http.StatusOK, nil
}
