package domain

import (
	"time"
)

type User struct {
	ID              string      `json:"_id,omitempty"`
	Email           string      `json:"email"`
	Password        string      `json:"password,omitempty"`
	PasswordConfirm string      `json:"passwordConfirm,omitempty"`
	Verified        bool        `json:"verified,omitempty"`
	VerifyCode      string      `json:"-"`
	UserRoles       []*UserRole `json:"userRoles,omitempty"`

	ApplicationCreatedAt time.Time `json:"applicationCreatedAt,omitempty"`
	CreatedAt            time.Time `json:"createdAt,omitempty"`
	UpdatedAt            time.Time `json:"updatedAt,omitempty"`
}