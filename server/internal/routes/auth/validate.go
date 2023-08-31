package auth

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

type SignUpInput struct {
	Email           string `json:"email"            validate:"required,email"`
	Password        string `json:"password"         validate:"required,min=8,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789,containsany=@!?"`
	PasswordConfirm string `json:"passwordConfirm"  validate:"required,min=8,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789,containsany=@!?"`
	Role            string `json:"role" validate:"required,oneof=user manager"`
}

func ValidateSignUpInput(inp *SignUpInput) error {
	validate := validator.New()
	err := validate.Struct(inp)

	if inp.Password != inp.PasswordConfirm {
		return ErrPasswordNotEqual
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("%s is required", err.Field()))
			case "email":
				return errors.New(fmt.Sprintf("%s is not a valid email", inp.Email))
			case "min":
				return errors.New(fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
			case "containsany":
				return errors.New(fmt.Sprintf("%s should contain at least one %s character", err.Field(), err.Param()))
			case "oneof":
				return errors.New(fmt.Sprintf("%s should include at least one of the following: %s", err.Field(), err.Param()))
			}
		}
	}

	return nil
}

type SendVerifyCodeInput struct {
	Email    string `json:"email"        validate:"required,email"`
	Password string `json:"password"     validate:"required,min=8,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789"`
}

func ValidateSendVerifyCodeInput(inp *SendVerifyCodeInput) error {
	validate := validator.New()
	err := validate.Struct(inp)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("%s is required", err.Field()))
			case "min":
				return errors.New(fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
			case "containsany":
				return errors.New(fmt.Sprintf("%s should contain at least one %s character", err.Field(), err.Param()))
			}
		}
	}

	return nil
}

type CheckVerifyCodeInput struct {
	Email      string `json:"email"        validate:"required,email"`
	Password   string `json:"password"     validate:"required,min=8,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789"`
	VerifyCode string `json:"verifyCode"   validate:"required,len=6"`
}

func ValidateCheckVerifyCodeInput(inp *CheckVerifyCodeInput) error {
	validate := validator.New()
	err := validate.Struct(inp)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("%s is required", err.Field()))
			case "min":
				return errors.New(fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
			case "containsany":
				return errors.New(fmt.Sprintf("%s should contain at least one %s character", err.Field(), err.Param()))
			case "len":
				return errors.New(fmt.Sprintf("%s must be %s characters", err.Field(), err.Param()))
			}
		}
	}

	return nil
}

type SignInInput struct {
	Email    string `json:"email"        validate:"required,email"`
	Password string `json:"password"     validate:"required,min=8,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789"`
}

func ValidateSignInInput(inp *SignInInput) error {
	validate := validator.New()
	err := validate.Struct(inp)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("%s is required", err.Field()))
			case "email":
				return errors.New(fmt.Sprintf("%s is not a valid email", inp.Email))
			case "min":
				return errors.New(fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
			case "containsany":
				return errors.New(fmt.Sprintf("%s should contain at least one %s character", err.Field(), err.Param()))
			}
		}
	}

	return nil
}

type Address struct {
	Country     string `json:"country" validate:"required"`
	City        string `json:"city" validate:"required"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
}

type GetProfileInput struct {
	ID    string `json:"_id,omitempty"`
	Email string `json:"email"`
}