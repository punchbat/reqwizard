package role

import (
	"context"
	"reqwizard/internal/domain"
)

type Repository interface {
	GetRoleByName(ctx context.Context, name string) (*domain.Role, error)
	GetRoleByID(ctx context.Context, id string) (*domain.Role, error)
	GetRoleByIDs(ctx context.Context, id []string) ([]*domain.Role, error)
	GetRoles(ctx context.Context) ([]*domain.Role, error)
}

type UseCase interface {
	GetRoles(ctx context.Context) ([]*domain.Role, int, error)
	GetRoleByID(ctx context.Context, id string) (*domain.Role, int, error)
}
