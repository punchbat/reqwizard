package handler

import (
	"net/http"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/auth"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	usecase auth.UseCase
}

func NewMiddleware(usecase auth.UseCase) gin.HandlerFunc {
	return (&Middleware{
		usecase: usecase,
	}).Handle
}

func (m *Middleware) Handle(c *gin.Context) {
	tokenFromHeader := c.GetHeader("Authorization")

	if tokenFromHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), tokenFromHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.BadResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})

		return
	}

	c.Set(auth.CtxUserKey, user)
}
