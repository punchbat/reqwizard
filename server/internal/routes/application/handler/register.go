package handler

import (
	"reqwizard/internal/routes/application/repository"
	"reqwizard/internal/routes/application/usecase"
	authRepo "reqwizard/internal/routes/auth/repository"
	"reqwizard/pkg/postgres/gorm"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, authMiddleware gin.HandlerFunc, isManagerMiddleware gin.HandlerFunc, db *gorm.Gorm) {
	// Создаем repository, все взаимодействия с db в ней
	repo := repository.NewRepository(db)
	authRepo := authRepo.NewRepository(db)

	// Создаем usecase, вся бизнес-логика в нем
	uc := usecase.NewUseCase(
		repo,
		authRepo,
	)

	// Create the handler
	h := NewHandler(uc)

	// Create the endpoints
	endpoints := router.Group("/application/v1")
	{
		endpoints.POST("/create", authMiddleware, h.CreateApplication)
		endpoints.GET("/download-file/:fileName", authMiddleware, h.GetFile)
		endpoints.GET("/:id", authMiddleware, h.GetApplicationByID)
		endpoints.GET("/my-list", authMiddleware, h.GetApplicationsByUserID)
		// * manager
		endpoints.GET("/manager-list", authMiddleware, isManagerMiddleware, h.GetApplications)
	}
}
