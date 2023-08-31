package ticketResponse

import (
	"errors"
	"fmt"
	"reqwizard/internal/shared/utils"

	"github.com/go-playground/validator"
)

type CreateTicketResponseInput struct {
	ID    string `json:"_id,omitempty"`
	Email string `json:"email"`

	ApplicationID string `json:"applicationId" validate:"required,len=36"`
	Text          string `json:"text" validate:"required,min=10"`
}

func ValidateCreateTicketResponseInput(inp *CreateTicketResponseInput) error {
	validate := validator.New()

	err := validate.Struct(inp)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return fmt.Errorf("%s is required", err.Field())
			case "len":
				return fmt.Errorf("%s must be exactly %s characters long", err.Field(), err.Param())
			case "min":
				return fmt.Errorf("%s should be at least %s characters long", err.Field(), err.Param())
			}
		}
	}

	return nil
}

type TicketResponseListInput struct {
	ID            string `form:"_id,omitempty"`
	Email         string `form:"email"`
	Search        string `form:"search,omitempty"`
	CreatedAtFrom string `form:"createdAtFrom,omitempty" validate:"omitempty,isodate_validation"`
	CreatedAtTo   string `form:"createdAtTo,omitempty" validate:"omitempty,isodate_validation"`
	UpdatedAtFrom string `form:"updatedAtFrom,omitempty" validate:"omitempty,isodate_validation"`
	UpdatedAtTo   string `form:"updatedAtTo,omitempty" validate:"omitempty,isodate_validation"`
}

func ValidateTicketResponseListInput(inp *TicketResponseListInput) error {
	validate := validator.New()
	validate.RegisterValidation("isodate_validation", utils.ValidateISODate)

	err := validate.Struct(inp)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "dive":
				return errors.New(fmt.Sprintf("Invalid value in %s", err.Field()))
			case "oneof":
				return errors.New(fmt.Sprintf("Invalid %s value(s): %s", err.Field(), err.Value()))
			case "isodate_validation":
				return errors.New(fmt.Sprintf("%s has an invalid date for the given type", err.Field()))
			}
		}
	}

	return nil
}

type TicketResponseInput struct {
	ID    string `json:"_id,omitempty"`
	Email string `json:"email"`
}