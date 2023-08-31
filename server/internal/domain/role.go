package domain

import (
	"time"
)

type RoleName string

const (
	RoleNameUser    RoleName = "user"
	RoleNameManager RoleName = "manager"
)

type Role struct {
	ID   string   `json:"_id,omitempty"`
	Name RoleName `json:"name,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}