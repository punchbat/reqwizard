package domain

import (
	"time"
)

type TicketResponse struct {
	ID            string `json:"_id,omitempty"`
	ApplicationID string `json:"applicationId,omitempty"`
	UserID        string `json:"userId,omitempty"`
	ManagerID     string `json:"managerId,omitempty"`
	Text          string `json:"text"`

	User    *User `json:"user,omitempty"`
	Manager *User `json:"manager,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}