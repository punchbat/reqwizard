package dto

import (
	"reqwizard/internal/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketResponse struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ApplicationID uuid.UUID `gorm:"not null"`
	UserID        uuid.UUID `gorm:"not null"`
	ManagerID     uuid.UUID
	Text          string `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func ConvertTicketResponseToDomain(i *TicketResponse) *domain.TicketResponse {
	return &domain.TicketResponse{
		ID:            i.ID.String(),
		ApplicationID: i.ApplicationID.String(),
		UserID:        i.UserID.String(),
		ManagerID:     i.ManagerID.String(),
		Text:          i.Text,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
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