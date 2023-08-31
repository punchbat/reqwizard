package userRole

import "errors"

var (
	ErrCantFindRole = errors.New("Can`t find role")
	ErrRoleIsExist = errors.New("This role already exists for the user")
	ErrRoleIsNotExist = errors.New("This role does not exist for the user")
)
