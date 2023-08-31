package usecase

import (
	"context"
	"errors"
	"reqwizard/internal/domain"

	"reqwizard/internal/routes/application"
	"reqwizard/internal/routes/auth"
	"reqwizard/internal/routes/ticketResponse"

	"github.com/google/uuid"
)

type UseCase struct {
	repo            ticketResponse.Repository
	applicationRepo application.Repository
	authRepo        auth.Repository
}

func NewUseCase(repo ticketResponse.Repository, applicationRepo application.Repository, authRepo auth.Repository) *UseCase {
	return &UseCase{
		repo:            repo,
		applicationRepo: applicationRepo,
		authRepo:        authRepo,
	}
}

func (a *UseCase) GetTicketResponseByID(ctx context.Context, id string) (*domain.TicketResponse, error) {
	ticketResponse, err := a.repo.GetTicketResponseByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ticketResponse, nil
}

func (a *UseCase) GetTicketResponsesByUserID(ctx context.Context, id string) ([]*domain.TicketResponse, error) {
	ticketResponses, err := a.repo.GetTicketResponsesByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ticketResponses, nil
}

// * manager.
func (a *UseCase) CreateTicketResponse(ctx context.Context, inp *ticketResponse.CreateTicketResponseInput) error {
	applcationEntity, err := a.applicationRepo.GetApplicationByID(ctx, inp.ApplicationID)
	if err != nil {
		return err
	}

	if len(applcationEntity.TicketResponseID) != 0 {
		return errors.New("You cannot create more than 1 ticket-response")
	}

	ticketResponseEntity := &domain.TicketResponse{
		ID:            uuid.New().String(),
		ApplicationID: applcationEntity.ID,
		UserID:        applcationEntity.UserID,
		ManagerID:     inp.ID,
		Text:          inp.Text,
	}

	err = a.repo.CreateTicketResponse(ctx, ticketResponseEntity)
	if err != nil {
		return err
	}

	applcationEntity.TicketResponseID = ticketResponseEntity.ID
	applcationEntity.ManagerID = inp.ID
	applcationEntity.Status = domain.ApplicationStatusDone
	_, err = a.applicationRepo.UpdateApplication(ctx, applcationEntity)
	if err != nil {
		return err
	}

	return nil
}

func (a *UseCase) GetTicketResponsesByManagerID(ctx context.Context, inp *ticketResponse.TicketResponseListInput) ([]*domain.TicketResponse, error) {
	ticketResponses, err := a.repo.GetTicketResponsesByManagerID(ctx, inp)
	if err != nil {
		return nil, err
	}

	return ticketResponses, nil
}