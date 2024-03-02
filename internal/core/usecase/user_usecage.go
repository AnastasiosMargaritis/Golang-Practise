package usecase

import (
	"context"
	"online-script/internal/core/domain"

	"github.com/jackc/pgx/v5/pgtype"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (userUseCase *userUseCase) CreateUser(ctx context.Context, arg domain.CreateUserReq) (domain.SecUser, error) {
	return userUseCase.userRepository.CreateUser(ctx, arg)
}

func (userUseCase *userUseCase) DeleteUser(ctx context.Context, userId pgtype.UUID) error {
	return userUseCase.userRepository.DeleteUser(ctx, userId)
}

func (userUseCase *userUseCase) GetUserById(ctx context.Context, userId pgtype.UUID) (domain.SecUser, error) {
	return userUseCase.userRepository.GetUserById(ctx, userId)
}

func (userUseCase *userUseCase) GetByUsername(ctx context.Context, username string) (domain.SecUser, error) {
	return userUseCase.userRepository.GetByUsername(ctx, username)
}

func (userUseCase *userUseCase) GetByEmail(ctx context.Context, email string) (domain.SecUser, error) {
	return userUseCase.userRepository.GetByEmail(ctx, email)
}
