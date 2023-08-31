package handler

import (
	"net/http"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/auth"
	"reqwizard/internal/routes/role"
	"reqwizard/internal/routes/userRole"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	roleName domain.RoleName
	status   domain.UserRoleStatus

	userRepo        auth.Repository
	roleRepo        role.Repository
	userRoleRepo    userRole.Repository
}

func NewMiddlewareUser(userRepo auth.Repository, roleRepo role.Repository, userRoleRepo userRole.Repository) gin.HandlerFunc {
	return (NewMiddleware(domain.RoleNameUser, domain.UserRoleStatusApproved, userRepo, roleRepo, userRoleRepo)).HandleMiddleware
}

func NewMiddlewareManager(userRepo auth.Repository, roleRepo role.Repository, userRoleRepo userRole.Repository) gin.HandlerFunc {
	return (NewMiddleware(domain.RoleNameManager, domain.UserRoleStatusApproved, userRepo, roleRepo, userRoleRepo)).HandleMiddleware
}

func NewMiddleware(roleName domain.RoleName, status domain.UserRoleStatus, userRepo auth.Repository, roleRepo role.Repository, userRoleRepo userRole.Repository) *Middleware {
	return &Middleware{roleName: roleName, status: status, userRepo: userRepo, roleRepo: roleRepo, userRoleRepo: userRoleRepo}
}

func (m *Middleware) HandleMiddleware(c *gin.Context) {
	user, exist := c.Get(auth.CtxUserKey)
	if !exist {
		c.JSON(http.StatusUnauthorized, domain.BadResponse{
			Status:  http.StatusUnauthorized,
			Message: auth.ErrUserIsUnauthorized.Error(),
		})
		c.Abort()
	}

	userEntity, err := m.userRepo.GetUserByID(c.Request.Context(), user.(*domain.User).ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.BadResponse{
			Status:  http.StatusUnauthorized,
			Message: auth.ErrUserNotFound.Error(),
		})
		c.Abort()
	}

	// * Нужно постоянно знать, что с ролями у пользователя
	userRoles, err := m.userRoleRepo.GetUserRoles(c.Request.Context(), userEntity.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.BadResponse{
			Status:  http.StatusUnauthorized,
			Message: role.ErrCantFindRole.Error(),
		})
		c.Abort()
	}

	for _, userRole := range userRoles {
		// * Нужно постоянно знать, что с ролями у пользователя
		roleEntity, err := m.roleRepo.GetRoleByID(c.Request.Context(), userRole.RoleID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.BadResponse{
				Status:  http.StatusUnauthorized,
				Message: role.ErrUserIsUnauthorized.Error(),
			})
			c.Abort()
		}

		isRole, isApproved := roleEntity.Name == m.roleName, userRole.Status == m.status

		if isRole && isApproved {
			c.Next()
			return
		}
	}

	c.JSON(http.StatusUnauthorized, domain.BadResponse{
		Status:  http.StatusUnauthorized,
		Message: role.ErrUserIsUnauthorized.Error(),
	})
	c.Abort()
}