package dto

type (
	CreatePostRequest struct {
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	CreatePostResponse struct {
		ID int64 `json:"id"`
	}
)
