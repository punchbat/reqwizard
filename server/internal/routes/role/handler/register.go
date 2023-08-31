package handler

import (
	authRepository "reqwizard/internal/routes/auth/repository"
	"reqwizard/internal/routes/role/repository"
	roleRepository "reqwizard/internal/routes/role/repository"
	"reqwizard/internal/routes/role/usecase"
	userRoleRepository "reqwizard/internal/routes/userRole/repository"
	"reqwizard/pkg/postgres/gorm"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, db *gorm.Gorm) (gin.HandlerFunc, gin.HandlerFunc) {
	// Создаем repository, все взаимодействия с db в ней
	repo := repository.NewRepository(db)
	authRepository := authRepository.NewRepository(db)
	roleRepository := roleRepository.NewRepository(db)
	userRoleRepository := userRoleRepository.NewRepository(db)

	// Создаем usecase, вся бизнес-логика в нем
	uc := usecase.NewUseCase(
		repo,
	)

	// Create the middleware instance
	isUserMiddleware := NewMiddlewareUser(authRepository, roleRepository, userRoleRepository)
	isManagerMiddleware := NewMiddlewareManager(authRepository, roleRepository, userRoleRepository)

	// Create the handler
	h := NewHandler(uc)

	// Create the endpoints
	endpoints := router.Group("/role/v1")
	{
		endpoints.GET("/list", h.GetRoles)
	}

	return isUserMiddleware, isManagerMiddleware
}
