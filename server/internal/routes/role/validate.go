package role

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

type RoleInput struct {
	ID    string `json:"_id,omitempty"`
	Email string `json:"email"`

	RoleID string `json:"roleId,omitempty" validate:"required"`
}

func ValidateRoleInput(inp *RoleInput) error {
	validate := validator.New()
	err := validate.Struct(inp)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("%s is required", err.Field()))
			}
		}
	}

	return nil
}