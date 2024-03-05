package usecase

import (
	"context"
	"online-script/internal/core/domain"
	"online-script/internal/util/tokenutil"
)

type LoginUsecase struct {
	userRepository domain.UserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository) domain.LoginUsecase {
	return &LoginUsecase{
		userRepository: userRepository,
	}
}

func (loginUsecase *LoginUsecase) GetUserByEmail(ctx context.Context, email string) (domain.SecUser, error) {
	return loginUsecase.userRepository.GetByEmail(ctx, email)
}

func (loginUsecase *LoginUsecase) CreateAccessToken(user *domain.SecUser, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (loginUsecase *LoginUsecase) CreateRefreshToken(user *domain.SecUser, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
