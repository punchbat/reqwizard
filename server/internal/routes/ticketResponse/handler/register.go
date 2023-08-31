package handler

import (
	applicationRepository "reqwizard/internal/routes/application/repository"
	authRepository "reqwizard/internal/routes/auth/repository"
	"reqwizard/internal/routes/ticketResponse/repository"
	"reqwizard/internal/routes/ticketResponse/usecase"
	"reqwizard/internal/services/email"
	"reqwizard/pkg/postgres/gorm"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, authMiddleware gin.HandlerFunc, isUserMiddleware gin.HandlerFunc, isManagerMiddleware gin.HandlerFunc, db *gorm.Gorm, mailer *email.Mailer) {
	// Создаем repository, все взаимодействия с db в ней
	repo := repository.NewRepository(db)
	applicationRepo := applicationRepository.NewRepository(db)
	authRepo := authRepository.NewRepository(db)

	// Создаем usecase, вся бизнес-логика в нем
	uc := usecase.NewUseCase(
		repo,
		applicationRepo,
		authRepo,

		mailer,
	)

	// Create the handler
	h := NewHandler(uc)

	// Create the endpoints
	endpoints := router.Group("/ticket-response/v1")
	{
		endpoints.GET("/:id", authMiddleware, isUserMiddleware, h.GetTicketResponseByID)
		endpoints.GET("/my-list", authMiddleware, isUserMiddleware, h.GetTicketResponsesByUserID)
		// * manager
		endpoints.POST("/create", authMiddleware, isManagerMiddleware, h.CreateTicketResponse)
		endpoints.GET("/manager-list", authMiddleware, isManagerMiddleware, h.GetTicketResponsesByManagerID)
	}
}
