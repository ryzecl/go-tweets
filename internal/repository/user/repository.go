package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, model *model.UserModel) (int64, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
