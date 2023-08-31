package role

import "errors"

var (
	ErrCantFindRole       = errors.New("Cant find role")
	ErrUserIsUnauthorized = errors.New("User is unauthorized")
)
