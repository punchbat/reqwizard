package utils

import "github.com/go-playground/validator"

func ValidateISODate(fl validator.FieldLevel) bool {
	date := fl.Field().String()

	if len(date) == 0 {
		return true
	}

	_, err := GetTimeFromString(date)
	if err != nil {
		return false
	}

	return true
}
