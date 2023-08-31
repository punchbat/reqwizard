package usecase

import (
	"context"
	"reqwizard/internal/domain"

	"reqwizard/internal/routes/role"
)

type UseCase struct {
	repoRole role.Repository
}

func NewUseCase(repoRole role.Repository) *UseCase {
	return &UseCase{
		repoRole: repoRole,
	}
}

func (uc *UseCase) GetRoles(ctx context.Context) ([]*domain.Role, error) {
	roles, err := uc.repoRole.GetRoles(ctx)

	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (uc *UseCase) GetRoleByID(ctx context.Context, id string) (*domain.Role, error) {
	role, err := uc.repoRole.GetRoleByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return role, nil
}