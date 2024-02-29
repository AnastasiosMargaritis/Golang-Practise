package domain

import "context"

type SecRole struct {
	RoleID   int64
	RoleName string
}

type RoleRepository interface {
	// Method fetches roles based on page and size
	ListRoles(ctx context.Context) ([]SecRole, error)
}

type RoleUsecase interface {
	// Method fetches roles based on page and size
	ListRoles(ctx context.Context) ([]SecRole, error)
}
