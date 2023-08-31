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

func (a *UseCase) GetRoles(ctx context.Context) ([]*domain.Role, error) {
	roles, err := a.repoRole.GetRoles(ctx)

	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (a *UseCase) GetRoleByID(ctx context.Context, id string) (*domain.Role, error) {
	role, err := a.repoRole.GetRoleByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return role, nil
}