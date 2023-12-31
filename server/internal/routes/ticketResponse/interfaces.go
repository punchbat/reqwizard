package ticketResponse

import (
	"context"
	"reqwizard/internal/domain"
)

type Repository interface {
	GetTicketResponseByID(ctx context.Context, id string) (*domain.TicketResponse, error)
	GetTicketResponsesByUserID(ctx context.Context, inp *TicketResponseListInput) ([]*domain.TicketResponse, error)

	CreateTicketResponse(ctx context.Context, req *domain.TicketResponse) error
	GetTicketResponses(ctx context.Context) ([]*domain.TicketResponse, error)
	GetTicketResponsesByManagerID(ctx context.Context, inp *TicketResponseListInput) ([]*domain.TicketResponse, error)
}

type UseCase interface {
	GetTicketResponseByID(ctx context.Context, id string) (*domain.TicketResponse, int, error)
	GetTicketResponsesByUserID(ctx context.Context, inp *TicketResponseListInput) ([]*domain.TicketResponse, int, error)

	// * manager
	CreateTicketResponse(ctx context.Context, inp *CreateTicketResponseInput) (int, error)
	GetTicketResponsesByManagerID(ctx context.Context, inp *TicketResponseListInput) ([]*domain.TicketResponse, int, error)
}
