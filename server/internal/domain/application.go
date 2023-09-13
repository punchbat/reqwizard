package domain

import (
	"time"
)

type ApplicationStatus string

const (
	ApplicationStatusCanceled ApplicationStatus = "canceled"
	ApplicationStatusWaiting  ApplicationStatus = "waiting"
	ApplicationStatusWorking  ApplicationStatus = "working"
	ApplicationStatusDone     ApplicationStatus = "done"
)

type ApplicationType string

const (
	ApplicationTypeGeneral ApplicationType = "general"

	ApplicationTypeFinancial ApplicationType = "financial"
)

type ApplicationSubType string

const (
	ApplicationSubTypeInformation ApplicationSubType = "information"
	ApplicationSubTypeAccountHelp ApplicationSubType = "account_help"

	ApplicationSubTypeRefunds ApplicationSubType = "refunds"
	ApplicationSubTypePayment ApplicationSubType = "payment"
)

type Application struct {
	ID               string             `json:"_id,omitempty"`
	TicketResponseID string             `json:"ticketResponseId,omitempty"`
	UserID           string             `json:"userId,omitempty"`
	ManagerID        string             `json:"managerId,omitempty"`
	Status           ApplicationStatus  `json:"status,omitempty"`
	Type             ApplicationType    `json:"type"`
	SubType          ApplicationSubType `json:"subType"`
	Title            string             `json:"title,omitempty"`
	Description      string             `json:"description,omitempty"`
	FileName         string             `json:"fileName,omitempty"`

	User    *User `json:"user,omitempty"`
	Manager *User `json:"manager,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}