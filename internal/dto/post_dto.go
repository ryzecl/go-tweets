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

type (
	Comment struct {
		ID        int64  `json:"id"`
		Username  string `json:"username"`
		Content   string `json:"content"`
		LikeCount int64  `json:"like_count"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	DetailPostResponse struct {
		ID        int64     `json:"id"`
		Username  string    `json:"username"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		LikeCount int64     `json:"like_count"`
		Comments  []Comment `json:"comments"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
	}
)
