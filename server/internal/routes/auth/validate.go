package auth

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reqwizard/internal/shared/utils"
	"time"

	"github.com/go-playground/validator"
)

type SignUpInput struct {
	Email           string         `form:"email"            validate:"required,email"`
	Password        string         `form:"password"         validate:"required,min=8,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789,containsany=@!?"`
	PasswordConfirm string         `form:"passwordConfirm"  validate:"required,min=8,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789,containsany=@!?"`
	Role            string         `form:"role" validate:"required,oneof=user manager"`
	Name            string         `form:"name" validate:"required,max=64"`
	Surname         string         `form:"surname" validate:"required,max=64"`
	Gender          string         `form:"gender" validate:"required,oneof=male female other"`
	Birthday        string         `form:"birthday" validate:"required,age_validation"`
	Avatar          multipart.File `form:"avatar,omitempty"`
	AvatarName string
}

func ValidateSignUpInput(inp *SignUpInput) error {
	validate := validator.New()
	validate.RegisterValidation("age_validation", validateAge)

	if inp.Password != inp.PasswordConfirm {
		return ErrPasswordNotEqual
	}

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
			case "max":
				return errors.New(fmt.Sprintf("%s cannot exceed %s characters", err.Field(), err.Param()))
			case "containsany":
				return errors.New(fmt.Sprintf("%s should contain at least one %s character", err.Field(), err.Param()))
			case "oneof":
				return errors.New(fmt.Sprintf("%s should include at least one of the following: %s", err.Field(), err.Param()))
			case "age_validation":
				return errors.New(fmt.Sprintf("%s age must not exceed 100 years", err.Field()))
			}
		}
	}

	return nil
}

func validateAge(fl validator.FieldLevel) bool {
	date := fl.Field().String()

	if len(date) == 0 {
		return true
	}

	dateTime, err := utils.GetTimeFromString(date)
	if err != nil {
		return false
	}

	today := time.Now()

	age := today.Year() - dateTime.Year()
	if today.YearDay() < dateTime.YearDay() {
		age--

		// Проверяем, что возраст не больше 100 лет и не в будущем
		return age <= 100 && age >= 10
	}

	return age <= 100
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