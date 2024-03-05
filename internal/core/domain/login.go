package domain

import "context"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (SecUser, error)
	CreateAccessToken(user *SecUser, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *SecUser, secret string, expiry int) (refreshToken string, err error)
}
