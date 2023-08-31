package handler

import (
	applicationRepository "reqwizard/internal/routes/application/repository"
	authRepository "reqwizard/internal/routes/auth/repository"
	"reqwizard/internal/routes/ticketResponse/repository"
	"reqwizard/internal/routes/ticketResponse/usecase"
	"reqwizard/pkg/postgres/gorm"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, authMiddleware gin.HandlerFunc, isManagerMiddleware gin.HandlerFunc, db *gorm.Gorm) {
	// Создаем repository, все взаимодействия с db в ней
	repo := repository.NewRepository(db)
	applicationRepo := applicationRepository.NewRepository(db)
	authRepo := authRepository.NewRepository(db)

	// Создаем usecase, вся бизнес-логика в нем
	uc := usecase.NewUseCase(
		repo,
		applicationRepo,
		authRepo,
	)

	// Create the handler
	h := NewHandler(uc)

	// Create the endpoints
	endpoints := router.Group("/ticket-response/v1")
	{
		endpoints.GET("/:id", authMiddleware, h.GetTicketResponseByID)
		endpoints.GET("/my-list", authMiddleware, h.GetTicketResponsesByUserID)
		// * manager
		endpoints.POST("/create", authMiddleware, isManagerMiddleware, h.CreateTicketResponse)
		endpoints.GET("/manager-list", authMiddleware, isManagerMiddleware, h.GetTicketResponsesByManagerID)
	}
}
