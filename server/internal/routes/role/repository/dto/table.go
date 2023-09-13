package dto

import (
	"reqwizard/internal/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	Name domain.RoleName `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func ConvertRoleToDomain(i *Role) *domain.Role {
	return &domain.Role{
		ID: i.ID.String(),

		Name: i.Name,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}

func ConvertRoleFromDomain(i *domain.Role) *Role {
	return &Role{
		ID: uuid.MustParse(i.ID),

		Name: i.Name,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}