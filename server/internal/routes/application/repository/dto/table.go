package dto

import (
	"database/sql"
	"reqwizard/internal/domain"
	authDto "reqwizard/internal/routes/auth/repository/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primarykey"`
	TicketResponseID *uuid.UUID
	UserID           uuid.UUID `gorm:"not null"`
	ManagerID        *uuid.UUID

	Status      domain.ApplicationStatus  `gorm:"not null"`
	Type        domain.ApplicationType    `gorm:"not null"`
	SubType     domain.ApplicationSubType `gorm:"not null"`
	Title       string                    `gorm:"not null"`
	Description string                    `gorm:"not null"`
	FileName    sql.NullString

	User    *authDto.User `gorm:"foreignKey:UserID;references:ID"`
	Manager *authDto.User `gorm:"foreignKey:ManagerID;references:ID"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func ConvertApplicationToDomain(i *Application) *domain.Application {
	application := &domain.Application{
		ID:          i.ID.String(),
		UserID:      i.UserID.String(),
		Status:      i.Status,
		Type:        i.Type,
		SubType:     i.SubType,
		Title:       i.Title,
		Description: i.Description,
		FileName:    i.FileName.String,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}

	if i.TicketResponseID != nil {
		application.TicketResponseID = i.TicketResponseID.String()
	}

	if i.ManagerID != nil {
		application.ManagerID = i.ManagerID.String()
	}

	if i.User != nil {
		application.User = authDto.ConvertUserToDomain(i.User)
	}

	if i.Manager != nil {
		application.Manager = authDto.ConvertUserToDomain(i.Manager)
	}

	return application
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