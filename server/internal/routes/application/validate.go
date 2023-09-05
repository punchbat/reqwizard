package application

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reqwizard/internal/shared/utils"

	"github.com/go-playground/validator"
)

type ApplicationInput struct {
	ID    string `json:"_id,omitempty"`
	Email string `json:"email"`
}

type CreateApplicationInput struct {
	ID    string `form:"_id,omitempty"`
	Email string `form:"email"`

	Type        string         `form:"type" validate:"required,oneof=financial general"`
	SubType     string         `form:"subType" validate:"required,subType_validation"`
	Title       string         `form:"title" validate:"required,min=10"`
	Description string         `form:"description" validate:"required,min=10"`
	File        multipart.File `form:"file"`
	FileName    string
}

func ValidateCreateApplicationInput(inp *CreateApplicationInput) error {
	validate := validator.New()
	validate.RegisterValidation("subType_validation", validateSubType)

	err := validate.Struct(inp)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("%s is required", err.Field()))
			case "oneof":
				return errors.New(fmt.Sprintf("%s must be either 'financial' or 'general'", err.Field()))
			case "min":
				return errors.New(fmt.Sprintf("%s should be at least %s characters long", err.Field(), err.Param()))
			case "subType_validation":
				return errors.New(fmt.Sprintf("%s has an invalid subType for the given type", err.Field()))
			}
		}
	}

	return nil
}

func validateSubType(fl validator.FieldLevel) bool {
	typeValue := fl.Parent().Elem().FieldByName("Type").String()
	subTypeValue := fl.Field().String()

	if typeValue == "financial" && (subTypeValue == "refunds" || subTypeValue == "payment") {
		return true
	} else if typeValue == "general" && (subTypeValue == "information" || subTypeValue == "account_help") {
		return true
	}

	return false
}

type ApplicationListInput struct {
	ID            string   `form:"_id,omitempty"`
	Email         string   `form:"email"`
	Search        string   `form:"search,omitempty"`
	Statuses      []string `form:"statuses,omitempty" binding:"dive" validate:"omitempty,dive,oneof=canceled waiting working done"`
	Types         []string `form:"types,omitempty" validate:"omitempty,dive,oneof=financial general"`
	SubTypes      []string `form:"subTypes,omitempty" validate:"omitempty,dive,oneof=refunds payment information account_help"`
	CreatedAtFrom string   `form:"createdAtFrom,omitempty" validate:"omitempty,isodate_validation"`
	CreatedAtTo   string   `form:"createdAtTo,omitempty" validate:"omitempty,isodate_validation"`
	UpdatedAtFrom string   `form:"updatedAtFrom,omitempty" validate:"omitempty,isodate_validation"`
	UpdatedAtTo   string   `form:"updatedAtTo,omitempty" validate:"omitempty,isodate_validation"`
}

func ValidateApplicationListInput(inp *ApplicationListInput) error {
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
			case "subType_validation":
				return errors.New(fmt.Sprintf("%s has an invalid subType for the given type", err.Field()))
			case "isodate_validation":
				return errors.New(fmt.Sprintf("%s has an invalid date for the given type", err.Field()))
			}
		}
	}

	return nil
}