package application

import (
	"context"
	"reqwizard/internal/domain"
)

type Repository interface {
	CreateApplication(ctx context.Context, req *domain.Application) error
	UpdateApplication(ctx context.Context, req *domain.Application) (*domain.Application, error)
	GetApplicationByID(ctx context.Context, id string) (*domain.Application, error)
	GetApplicationsByUserID(ctx context.Context, inp *ApplicationListInput) ([]*domain.Application, error)
	GetApplications(ctx context.Context, inp *ApplicationListInput) ([]*domain.Application, error)
}

type UseCase interface {
	CreateApplication(ctx context.Context, inp *CreateApplicationInput) (int, error)
	GetFile(ctx context.Context, userID string, fileName string) ([]byte, string, int, error)
	GetApplicationByID(ctx context.Context, id string) (*domain.Application, int, error)
	GetApplicationsByUserID(ctx context.Context, inp *ApplicationListInput) ([]*domain.Application, int, error)

	// * manager
	GetApplications(ctx context.Context, inp *ApplicationListInput) ([]*domain.Application, int, error)
}
