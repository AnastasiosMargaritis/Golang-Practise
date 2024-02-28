package usecase

import (
	"context"
	"online-script/internal/core/domain"
)

type RoleUsecase struct {
	roleRepository domain.RoleRepository
}

func NewRoleUseCase(roleRepository domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		roleRepository: roleRepository,
	}
}

func (roleUsecase *RoleUsecase) ListRoles(c context.Context, params domain.ListRolesParams) ([]domain.SecRole, error) {
	return roleUsecase.roleRepository.ListRoles(c, params)
}
