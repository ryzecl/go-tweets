package user

import (
	"context"
	"go-tweets/internal/config"
	"go-tweets/internal/dto"
	"go-tweets/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int64) (string, string, int, error)
}

type userService struct {
	cfg      *config.Config
	userRepo user.UserRepository
}

func NewUserService(cfg *config.Config, userRepo user.UserRepository) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}
