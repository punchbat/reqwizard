package dto

import (
	"reqwizard/internal/domain"
	authDto "reqwizard/internal/routes/auth/repository/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketResponse struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ApplicationID uuid.UUID `gorm:"not null"`
	UserID        uuid.UUID `gorm:"not null"`
	ManagerID     uuid.UUID
	Text          string `gorm:"not null"`

	User    *authDto.User `gorm:"foreignKey:UserID;references:ID"`
	Manager *authDto.User `gorm:"foreignKey:ManagerID;references:ID"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func ConvertTicketResponseToDomain(i *TicketResponse) *domain.TicketResponse {
	ticketResponse := &domain.TicketResponse{
		ID:            i.ID.String(),
		ApplicationID: i.ApplicationID.String(),
		UserID:        i.UserID.String(),
		ManagerID:     i.ManagerID.String(),
		Text:          i.Text,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}

	if i.User != nil {
		ticketResponse.User = authDto.ConvertUserToDomain(i.User)
	}

	if i.Manager != nil {
		ticketResponse.Manager = authDto.ConvertUserToDomain(i.Manager)
	}

	return ticketResponse
}

func ConvertTicketResponseFromDomain(i *domain.TicketResponse) *TicketResponse {
	return &TicketResponse{
		ID:            uuid.MustParse(i.ID),
		ApplicationID: uuid.MustParse(i.ApplicationID),
		UserID:        uuid.MustParse(i.UserID),
		ManagerID:     uuid.MustParse(i.ManagerID),
		Text:          i.Text,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}