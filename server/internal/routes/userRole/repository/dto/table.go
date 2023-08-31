package dto

import (
	"reqwizard/internal/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID string    `gorm:"not null"`
	RoleID string    `gorm:"not null"`
	Status string    `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func ConvertUserRoleToDomain(i *UserRole) *domain.UserRole {
	return &domain.UserRole{
		ID:     i.ID.String(),
		UserID: i.UserID,
		RoleID: i.RoleID,
		Status: domain.UserRoleStatus(i.Status),

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}

func ConvertUserRoleFromDomain(i *domain.UserRole) *UserRole {
	return &UserRole{
		ID:     uuid.MustParse(i.ID),
		UserID: i.UserID,
		RoleID: i.RoleID,
		Status: string(i.Status),

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}