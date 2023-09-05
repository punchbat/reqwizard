package domain

import (
	"time"
)

type UserGender string

const (
	UserGenderMale   UserGender = "male"
	UserGenderFemale UserGender = "female"
	UserGenderOther  UserGender = "other"
)

type User struct {
	ID              string      `json:"_id,omitempty"`
	Email           string      `json:"email"`
	Password        string      `json:"password,omitempty"`
	PasswordConfirm string      `json:"passwordConfirm,omitempty"`
	Verified        bool        `json:"verified,omitempty"`
	VerifyCode      string      `json:"-"`
	UserRoles       []*UserRole `json:"userRoles,omitempty"`

	Name     string     `json:"name,omitempty"`
	Surname  string     `json:"surname,omitempty"`
	Gender   UserGender `json:"gender,omitempty"`
	Birthday time.Time  `json:"birthday,omitempty"`
	Avatar   string     `json:"avatar,omitempty"`

	ApplicationCreatedAt time.Time `json:"applicationCreatedAt,omitempty"`
	CreatedAt            time.Time `json:"createdAt,omitempty"`
	UpdatedAt            time.Time `json:"updatedAt,omitempty"`
}