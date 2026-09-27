package model

import "time"

type (
	PostModel struct {
		ID        int64
		UserID    int64
		Title     string
		Content   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	}

	PostLikeModel struct {
		ID        int64
		PostID    int64
		UserID    int64
		CreatedAt time.Time
	}

	PostWithUserModel struct {
		ID        int64
		UserID    int64
		Username  string
		Title     string
		Content   string
		LikeCount int64
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	}
)
