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
)

func (s *userService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int64) (string, string, int, error) {
	// check user exists
	userExists, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExists == nil {
		return "", "", http.StatusNotFound, errors.New("User not found")
	}

	// get refresh token by user id
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExists == nil {
		return "", "", http.StatusUnauthorized, errors.New("Refresh token was expired")
	}

	// check refresh token is match with request body
	if refreshTokenExists.RefreshToken != req.RefreshToken {
		return "", "", http.StatusUnauthorized, errors.New("Refresh token is invalid")
	}

	// generate new token
	token, err := jwt.CreateToken(userID, userExists.Username, s.cfg.SecretJWT)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// delete old refresh token and generate new one
	err = s.userRepo.DeleteRefreshTokenByUserID(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	now := time.Now()
	s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userID,
		RefreshToken: refreshToken,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	})

	return token, refreshToken, http.StatusOK, nil
}
