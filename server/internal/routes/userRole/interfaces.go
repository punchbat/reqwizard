package userRole

import (
	"context"
	"reqwizard/internal/domain"
)

type Repository interface {
	CreateUserRole(ctx context.Context, userRole *domain.UserRole) error
	GetUserRoleByID(ctx context.Context, id string) (*domain.UserRole, error)
	GetUserRoles(ctx context.Context, userId string) ([]*domain.UserRole, error)
	DeleteUserRoleByID(ctx context.Context, id string) error
	UpdateUserRoles(ctx context.Context, userRoles []*domain.UserRole) error
}