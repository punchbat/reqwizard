package handler

import (
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/role"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase role.UseCase
}

func NewHandler(useCase role.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// GetRoles
// @Tags roles
// @Summary receiving roles
// @Description get all roles
// @Success 200 {object} domain.ResponseRoles
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/role/v1/list [get].
func (h *Handler) GetRoles(c *gin.Context) {
	roles, status, err := h.useCase.GetRoles(c.Request.Context())
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})
		return
	}

	c.JSON(status, domain.ResponseRoles{
		Status:  status,
		Payload: roles,
	})
}
