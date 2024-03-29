package usecase

import (
	"context"
	"online-script/internal/core/domain"
)

type roleUsecase struct {
	roleRepository domain.RoleRepository
}

func NewRoleUseCase(roleRepository domain.RoleRepository) domain.RoleUsecase {
	return &roleUsecase{
		roleRepository: roleRepository,
	}
}

func (roleUsecase *roleUsecase) ListRoles(c context.Context) ([]domain.SecRole, error) {
	return roleUsecase.roleRepository.ListRoles(c)
}

func (roleUsecase *roleUsecase) GetRole(ctx context.Context, roleID int64) (domain.SecRole, error) {
	return roleUsecase.roleRepository.GetRole(ctx, roleID)
}
