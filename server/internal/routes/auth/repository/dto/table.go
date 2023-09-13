package dto

import (
	"database/sql"
	"reqwizard/internal/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	Email           string `gorm:"not null"`
	Password        string `gorm:"not null"`
	PasswordConfirm string `gorm:"not null"`
	Verified        bool   `gorm:"not null"`
	VerifyCode      sql.NullString

	Name     string            `gorm:"type:not null;varchar(64)"`
	Surname  string            `gorm:"type:not null;varchar(64)"`
	Gender   domain.UserGender `gorm:"type:not null;varchar(6)"`
	Birthday time.Time         `gorm:"not null"`
	Avatar   sql.NullString

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

		Name:     i.Name,
		Surname:  i.Surname,
		Gender:   i.Gender,
		Birthday: i.Birthday,
		Avatar:   i.Avatar.String,

		ApplicationCreatedAt: i.ApplicationCreatedAt.Time,
		CreatedAt:            i.CreatedAt,
		UpdatedAt:            i.UpdatedAt,
	}
}

func ConvertUserFromDomain(i *domain.User) *User {
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
		VerifyCode:      sql.NullString{String: i.VerifyCode, Valid: true},

		Name:     i.Name,
		Surname:  i.Surname,
		Gender:   i.Gender,
		Birthday: i.Birthday,
		Avatar:   sql.NullString{String: i.Avatar, Valid: true},

		ApplicationCreatedAt: applicationCreatedAt,
		CreatedAt:            i.CreatedAt,
		UpdatedAt:            i.UpdatedAt,
	}
}