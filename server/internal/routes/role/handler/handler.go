package handler

import (
	"net/http"
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
	roles, err := h.useCase.GetRoles(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.ResponseRoles{
		Status:  http.StatusOK,
		Payload: roles,
	})
}
