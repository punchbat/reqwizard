package usecase

import (
	"context"
	"net/http"
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

func (uc *UseCase) GetRoles(ctx context.Context) ([]*domain.Role, int, error) {
	roles, err := uc.repoRole.GetRoles(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return roles, http.StatusOK, nil
}

func (uc *UseCase) GetRoleByID(ctx context.Context, id string) (*domain.Role, int, error) {
	role, err := uc.repoRole.GetRoleByID(ctx, id)

	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return role, http.StatusOK, nil
}