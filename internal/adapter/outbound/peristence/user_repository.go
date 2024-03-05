package persistence

import (
	"context"
	"fmt"

	"online-script/internal/core/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// RoleRepository implements RoleRepository interface using pgx.
type UserRepository struct {
	db *pgx.Conn // or *pgx.ConnPool if you're using a connection pool
}

// NewRoleRepository creates a new instance of RoleRepository.
func NewUserRepository(db *pgx.Conn) domain.UserRepository {
	return &UserRepository{db: db}
}

const createUser = `-- name: CreateUser :one
INSERT INTO public.sec_user
(user_id, username, hashed_password, first_name, last_name, email, role_id)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING user_id, username, hashed_password, first_name, last_name, email, created_at, role_id
`

type CreateUserParams struct {
	UserID         pgtype.UUID
	Username       string
	HashedPassword string
	FirstName      string
	LastName       string
	Email          string
	RoleID         int64
}

func (userRepository *UserRepository) CreateUser(ctx context.Context, req domain.CreateUserReq) (domain.SecUser, error) {
	userId := uuid.New()
	arg := CreateUserParams{
		UserID:         pgtype.UUID{Bytes: userId, Valid: true},
		Username:       req.Username,
		HashedPassword: req.Password,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		RoleID:         req.RoleID,
	}

	fmt.Println(arg)
	row := userRepository.db.QueryRow(ctx, createUser,
		arg.UserID,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.RoleID,
	)
	var i domain.SecUser
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.RoleID,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM public.sec_user where user_id = $1
`

func (userRepository *UserRepository) DeleteUser(ctx context.Context, userID pgtype.UUID) error {
	_, err := userRepository.db.Exec(ctx, deleteUser, userID)
	return err
}

const getUserById = `-- name: GetUserById :one
SELECT user_id, username, hashed_password, first_name, last_name, email, created_at, role_id FROM public.sec_user
WHERE user_id = $1
`

func (userRepository *UserRepository) GetUserById(ctx context.Context, userID pgtype.UUID) (domain.SecUser, error) {
	row := userRepository.db.QueryRow(ctx, getUserById, userID)
	var i domain.SecUser
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.RoleID,
	)
	return i, err
}

const getByEmail = `-- name: GetByEmail :one
SELECT user_id, username, hashed_password, first_name, last_name, email, created_at, role_id FROM public.sec_user
WHERE email = $1
`

func (userRepository *UserRepository) GetByEmail(ctx context.Context, email string) (domain.SecUser, error) {
	row := userRepository.db.QueryRow(ctx, getByEmail, email)
	var i domain.SecUser
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.RoleID,
	)
	return i, err
}

const getByUsername = `-- name: GetByUsername :one
SELECT user_id, username, hashed_password, first_name, last_name, email, created_at, role_id FROM public.sec_user
WHERE username = $1
`

func (userRepository *UserRepository) GetByUsername(ctx context.Context, username string) (domain.SecUser, error) {
	row := userRepository.db.QueryRow(ctx, getByUsername, username)
	var i domain.SecUser
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.RoleID,
	)
	return i, err
}
