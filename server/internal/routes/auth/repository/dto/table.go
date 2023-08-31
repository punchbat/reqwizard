package dto

import (
	"database/sql"
	"reqwizard/internal/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`

	Email           string `gorm:"not null"`
	Password        string `gorm:"not null"`
	PasswordConfirm string `gorm:"not null"`
	Verified        bool   `gorm:"not null"`
	VerifyCode      sql.NullString

	ApplicationCreatedAt sql.NullTime
	CreatedAt            time.Time `gorm:"not null"`
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt
}

func ConvertUserToDomain(i *User) *domain.User {
	return &domain.User{
		ID:              i.ID.String(),
		Email:           i.Email,
		Password:        i.Password,
		PasswordConfirm: i.PasswordConfirm,
		Verified:        i.Verified,
		VerifyCode:      i.VerifyCode.String,

		ApplicationCreatedAt: i.ApplicationCreatedAt.Time,
		CreatedAt:            i.CreatedAt,
		UpdatedAt:            i.UpdatedAt,
	}
}

func ConvertUserFromDomain(i *domain.User) *User {
	var verifyCode sql.NullString
	if i.VerifyCode != "" {
		verifyCode.String = i.VerifyCode
		verifyCode.Valid = true
	}

	var applicationCreatedAt sql.NullTime
	if !i.ApplicationCreatedAt.IsZero() {
		applicationCreatedAt.Time = i.ApplicationCreatedAt
		applicationCreatedAt.Valid = true
	}

	return &User{
		ID:              uuid.MustParse(i.ID),
		Email:           i.Email,
		Password:        i.Password,
		PasswordConfirm: i.PasswordConfirm,
		Verified:        i.Verified,
		VerifyCode:      verifyCode,

		ApplicationCreatedAt: applicationCreatedAt,
		CreatedAt:            i.CreatedAt,
		UpdatedAt:            i.UpdatedAt,
	}
}