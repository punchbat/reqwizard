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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(router *gin.Engine, pgGorm *gorm.Gorm, mailer *email.Mailer) {
	//* Свагер
	docs.SwaggerInfo.BasePath = "/api"
	
	//* Пингуем сервер
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// * AUTH
	authMiddleware := authHandler.RegisterHTTPEndpoints(router, pgGorm, mailer)

	// * API endpoints
	api := router.Group("/api")

	// * ROLE
	_, isManagerMiddleware := roleHandler.RegisterHTTPEndpoints(api, authMiddleware, pgGorm)

	// * APPLICATION
	applicationHandler.RegisterHTTPEndpoints(api, authMiddleware, isManagerMiddleware, pgGorm)
	
	// * TICKET_RESPONSE
	ticketResponseHandler.RegisterHTTPEndpoints(api, authMiddleware, isManagerMiddleware, pgGorm)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}