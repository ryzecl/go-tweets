package dto

type (
	CreateOrUpdatePostRequest struct {
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	CreateOrUpdatePostResponse struct {
		ID int64 `json:"id"`
	}
)

type (
	LikeOrUnlikePostRequest struct {
		PostID int64 `json:"post_id" validate:"required"`
	}
)
