package usecase

import (
	"context"
	"errors"
	"fmt"
	"reqwizard/internal/domain"
	"reqwizard/internal/services/email"

	"reqwizard/internal/routes/application"
	"reqwizard/internal/routes/auth"
	"reqwizard/internal/routes/ticketResponse"

	service_email "reqwizard/internal/services/email"

	"github.com/google/uuid"
)

type UseCase struct {
	repo            ticketResponse.Repository
	applicationRepo application.Repository
	authRepo        auth.Repository

	mailer *email.Mailer
}

func NewUseCase(repo ticketResponse.Repository, applicationRepo application.Repository, authRepo auth.Repository,

	mailer *service_email.Mailer,
) *UseCase {
	return &UseCase{
		repo:            repo,
		applicationRepo: applicationRepo,
		authRepo:        authRepo,

		mailer: mailer,
	}
}

func (uc *UseCase) GetTicketResponseByID(ctx context.Context, id string) (*domain.TicketResponse, error) {
	ticketResponse, err := uc.repo.GetTicketResponseByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ticketResponse, nil
}

func (uc *UseCase) GetTicketResponsesByUserID(ctx context.Context, inp *ticketResponse.TicketResponseListInput) ([]*domain.TicketResponse, error) {
	ticketResponses, err := uc.repo.GetTicketResponsesByUserID(ctx, inp)
	if err != nil {
		return nil, err
	}

	return ticketResponses, nil
}

// * manager.

type EmailContent struct {
	ApplicationTitle   string
	TicketResponseText string
	Link               string
}

func (uc *UseCase) CreateTicketResponse(ctx context.Context, inp *ticketResponse.CreateTicketResponseInput) error {
	applcationEntity, err := uc.applicationRepo.GetApplicationByID(ctx, inp.ApplicationID)
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

	err = uc.repo.CreateTicketResponse(ctx, ticketResponseEntity)
	if err != nil {
		return err
	}

	applcationEntity.TicketResponseID = ticketResponseEntity.ID
	applcationEntity.ManagerID = inp.ID
	applcationEntity.Status = domain.ApplicationStatusDone
	_, err = uc.applicationRepo.UpdateApplication(ctx, applcationEntity)
	if err != nil {
		return err
	}

	userEntity, err := uc.authRepo.GetUserByID(ctx, applcationEntity.UserID)
	if err != nil {
		return err
	}

	// Отправляем письмо
	emailMessage := service_email.Message{
		Subject:      "Reqwizard: your Application is done",
		To:           []string{userEntity.Email},
		TemplateName: "UserApplicationDone",
		Content: EmailContent{
			ApplicationTitle:   applcationEntity.Title,
			TicketResponseText: ticketResponseEntity.Text,
			Link:               fmt.Sprintf("http://localhost:8000/ticket-response/%s", ticketResponseEntity.ID),
		},
	}
	uc.mailer.Send(&emailMessage)

	return nil
}

func (uc *UseCase) GetTicketResponsesByManagerID(ctx context.Context, inp *ticketResponse.TicketResponseListInput) ([]*domain.TicketResponse, error) {
	ticketResponses, err := uc.repo.GetTicketResponsesByManagerID(ctx, inp)
	if err != nil {
		return nil, err
	}

	return ticketResponses, nil
}