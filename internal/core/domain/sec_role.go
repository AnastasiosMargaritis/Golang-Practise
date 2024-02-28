package domain

import "context"

type SecRole struct {
	RoleID   int64
	RoleName string
}

type ListRolesParams struct {
	Limit  int32
	Offset int32
}

type RoleRepository interface {
	ListRoles(ctx context.Context, params ListRolesParams) ([]SecRole, error)
}

type RoleUsecase interface {
	ListRoles(ctx context.Context, params ListRolesParams) ([]SecRole, error)
}
