package handler

import (
	"net/http"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/auth"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	tokenFromCookie, err := c.Cookie(viper.GetString("auth.token.name"))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, status, err := m.usecase.ParseToken(c.Request.Context(), tokenFromCookie)
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})

		return
	}

	c.Set(auth.CtxUserKey, user)
}
