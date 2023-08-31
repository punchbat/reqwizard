package utils

import (
	"reqwizard/internal/domain"
)

func IsUser(slice []*domain.UserRole) bool {
	index, _ := Find(slice, func(userRole *domain.UserRole) bool { return userRole.Name == domain.UserRoleNameUser })

	return index != -1
}

func IsManager(slice []*domain.UserRole) bool {
	index, _ := Find(slice, func(userRole *domain.UserRole) bool { return userRole.Name == domain.UserRoleNameManager })

	return index != -1
}