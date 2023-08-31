package auth

import "errors"

var (
	ErrEmailOrPassword         = errors.New("Email or Password is wrong")
	ErrPasswordNotEqual        = errors.New("Password not equal")
	ErrUserIsUnauthorized      = errors.New("User is unauthorized")
	ErrUserNotFound            = errors.New("User not found")
	ErrUserIsExist             = errors.New("User is exist")
	ErrVerifyCodeNotMatch      = errors.New("Verify code is invalid")
	ErrInvalidAccessToken      = errors.New("Invalid access token")
	ErrRoleIsExist             = errors.New("This role already exists for the user")
	ErrCantFindUserRole        = errors.New("Can`t find role")
	ErrCantDeleteUserRole      = errors.New("Can`t delete role")
	ErrCantUpdateUser          = errors.New("Can`t update user")
	ErrNotFoundUserRoleForUser = errors.New("Not found user role for the user")
)
