package dto

import (
	"database/sql"
	"reqwizard/internal/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	TicketResponseID *uuid.UUID
	UserID           uuid.UUID `gorm:"not null"`
	ManagerID        *uuid.UUID

	Status      domain.ApplicationStatus  `gorm:"not null"`
	Type        domain.ApplicationType    `gorm:"not null"`
	SubType     domain.ApplicationSubType `gorm:"not null"`
	Title       string                    `gorm:"not null"`
	Description string                    `gorm:"not null"`
	FileName    sql.NullString

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func ConvertApplicationToDomain(i *Application) *domain.Application {
	var managerID string
	if i.ManagerID != nil {
		managerID = i.ManagerID.String()
	}

	var ticketResponseID string
	if i.TicketResponseID != nil {
		ticketResponseID = i.TicketResponseID.String()
	}

	return &domain.Application{
		ID:               i.ID.String(),
		TicketResponseID: ticketResponseID,
		UserID:           i.UserID.String(),
		ManagerID:        managerID,
		Status:           i.Status,
		Type:             i.Type,
		SubType:          i.SubType,
		Title:            i.Title,
		Description:      i.Description,
		FileName:         i.FileName.String,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}

func ConvertApplicationFromDomain(i *domain.Application) *Application {
	application := &Application{
		ID:          uuid.MustParse(i.ID),
		UserID:      uuid.MustParse(i.UserID),
		Status:      i.Status,
		Type:        i.Type,
		SubType:     i.SubType,
		Title:       i.Title,
		Description: i.Description,
		FileName:    sql.NullString{String: i.FileName, Valid: true},

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}

	if i.ManagerID != "" {
		managerID := uuid.MustParse(i.ManagerID)
		application.ManagerID = &managerID
	}

	if i.TicketResponseID != "" {
		ticketResponseID := uuid.MustParse(i.TicketResponseID)
		application.TicketResponseID = &ticketResponseID
	}

	return application
}