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
)
