package domain

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type SecUser struct {
	UserID         pgtype.UUID
	Username       string
	HashedPassword string
	FirstName      string
	LastName       string
	Email          string
	CreatedAt      pgtype.Timestamptz
	RoleID         pgtype.Int8
}

type CreateUserReq struct {
	Username  string `json:"username" binding:"required,alphanum"`
	Password  string `json:"password" binding:"required,alphanum"`
	FirstName string `json:"firstName" binding:"required,alphanum"`
	LastName  string `json:"lastName" binding:"required,alphanum"`
	Email     string `json:"email" binding:"required,email"`
	RoleID    int64  `json:"roleId" binding:"required,min=1"`
}

type UserRepository interface {
	// Method to create a new User
	CreateUser(ctx context.Context, arg CreateUserReq) (SecUser, error)

	// Method to delete a user by id
	DeleteUser(ctx context.Context, userId pgtype.UUID) error

	// Find a user by Id
	GetUserById(ctx context.Context, userId pgtype.UUID) (SecUser, error)
}

type UserUseCase interface {
	// Method to create a new User
	CreateUser(ctx context.Context, arg CreateUserReq) (SecUser, error)

	// Method to delete a user by id
	DeleteUser(ctx context.Context, userId pgtype.UUID) error

	// Find a user by Id
	GetUserById(ctx context.Context, userId pgtype.UUID) (SecUser, error)
}
