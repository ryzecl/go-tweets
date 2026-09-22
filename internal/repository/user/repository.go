package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
	"time"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, model *model.UserModel) (int64, error)
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*model.RefreshTokenModel, error)
	StoreRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error
	GetUserByID(ctx context.Context, userID int64) (*model.UserModel, error)
	DeleteRefreshTokenByUserID(ctx context.Context, userID int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
