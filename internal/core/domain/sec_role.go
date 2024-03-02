package domain

import "context"

type SecRole struct {
	RoleID   int64
	RoleName string
}

type RoleRepository interface {
	// Method fetches all roles
	ListRoles(ctx context.Context) ([]SecRole, error)

	// Method fetches role by id
	GetRole(ctx context.Context, roleID int64) (SecRole, error)
}

type RoleUsecase interface {
	// Method fetches roles based on page and size
	ListRoles(ctx context.Context) ([]SecRole, error)

	// Method fetches roles based on page and size
	GetRole(ctx context.Context, roleID int64) (SecRole, error)
}
