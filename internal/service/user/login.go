package user

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"go-tweets/pkg/jwt"
	"go-tweets/pkg/refreshtoken"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
	// Check user exists
	userExists, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExists == nil {
		return "", "", http.StatusUnauthorized, errors.New("Wrong email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(req.Password))
	if err != nil {
		return "", "", http.StatusUnauthorized, errors.New("Wrong email or password")
	}

	// Generate access token
	token, err := jwt.CreateToken(userExists.ID, userExists.Username, s.cfg.SecretJWT)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// Get refresh token if exists
	now := time.Now()
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userExists.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExists != nil {
		return token, refreshTokenExists.RefreshToken, http.StatusOK, nil
	}

	// Generate and store refresh token
	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userExists.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return token, refreshToken, http.StatusOK, nil
}
