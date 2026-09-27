package dto

type (
	StoreCommentRequest struct {
		PostID  int64  `json:"post_id" validate:"required"`
		Content string `json:"content" validate:"required"`
	}
)
