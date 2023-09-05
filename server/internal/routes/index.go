package routes

import (
	"net/http"
	"reqwizard/docs"

	applicationHandler "reqwizard/internal/routes/application/handler"
	authHandler "reqwizard/internal/routes/auth/handler"
	roleHandler "reqwizard/internal/routes/role/handler"
	ticketResponseHandler "reqwizard/internal/routes/ticketResponse/handler"
	"reqwizard/internal/services/email"
	"reqwizard/pkg/postgres/gorm"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(router *gin.Engine, c *cron.Cron, pgGorm *gorm.Gorm, mailer *email.Mailer) {
	// * Свагер
	docs.SwaggerInfo.BasePath = "/api"

	// * Пингуем сервер
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// * раздача статики аватарок
	router.Static("/uploads/avatars", "./uploads/avatars")

	// * AUTH
	authMiddleware := authHandler.RegisterHTTPEndpoints(router, c, pgGorm, mailer)

	// * API endpoints
	api := router.Group("/api")

	// * ROLE
	isUserMiddleware, isManagerMiddleware := roleHandler.RegisterHTTPEndpoints(api, pgGorm)

	// * APPLICATION
	applicationHandler.RegisterHTTPEndpoints(api, authMiddleware, isUserMiddleware, isManagerMiddleware, pgGorm)

	// * TICKET_RESPONSE
	ticketResponseHandler.RegisterHTTPEndpoints(api, authMiddleware, isUserMiddleware, isManagerMiddleware, pgGorm, mailer)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}