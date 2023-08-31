package domain

import (
	"time"
)

type UserRoleName string

const (
	UserRoleNameUser    UserRoleName = "user"
	UserRoleNameManager UserRoleName = "manager"
)

type UserRoleStatus string

const (
	UserRoleStatusCanceled UserRoleStatus = "canceled"
	UserRoleStatusPending  UserRoleStatus = "pending"
	UserRoleStatusApproved UserRoleStatus = "approved"
)

type UserRole struct {
	ID     string         `json:"_id,omitempty"`
	UserID string         `json:"userId,omitempty"`
	RoleID string         `json:"roleId,omitempty"`
	Name   UserRoleName   `json:"name,omitempty"`
	Status UserRoleStatus `json:"status,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}