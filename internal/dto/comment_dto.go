package dto

type (
	StoreCommentRequest struct {
		PostID  int64  `json:"post_id" validate:"required"`
		Content string `json:"content" validate:"required"`
	}
)

type (
	LikeOrUnlikeCommentRequest struct {
		CommentID int64 `json:"comment_id" validate:"required"`
	}
)
