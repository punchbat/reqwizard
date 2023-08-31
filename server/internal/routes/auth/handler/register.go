package handler

import (
	"reqwizard/internal/routes/auth/jobs"
	"reqwizard/internal/routes/auth/repository"
	"reqwizard/internal/routes/auth/usecase"
	roleRepository "reqwizard/internal/routes/role/repository"
	userRoleRepository "reqwizard/internal/routes/userRole/repository"
	"reqwizard/internal/services/email"
	"reqwizard/pkg/postgres/gorm"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

func RegisterHTTPEndpoints(router *gin.Engine, c *cron.Cron, db *gorm.Gorm, mailer *email.Mailer) gin.HandlerFunc {
	// Создаем repository, все взаимодействия с db в ней
	repo := repository.NewRepository(db)
	roleRepository := roleRepository.NewRepository(db)
	userRoleRepository := userRoleRepository.NewRepository(db)

	// Создаем usecase, вся бизнес-логика в нем
	uc := usecase.NewUseCase(
		repo,
		roleRepository,
		userRoleRepository,

		mailer,
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl"),
	)

	// Jobs
	authJobScheduler := jobs.NewAuthJobScheduler(uc)
	authJobScheduler.Start(c)

	// Create the middleware instance
	authMiddleware := NewMiddleware(uc)

	// Create the handler
	h := NewHandler(uc)

	// Create the endpoints
	endpoints := router.Group("/auth/v1")
	{
		endpoints.POST("/sign-up", h.SignUp)
		endpoints.POST("/send-verify-code", h.SendVerifyCode)
		endpoints.POST("/check-verify-code", h.CheckVerifyCode)
		endpoints.POST("/sign-in", h.SignIn)

		// * проверяем на наличие аутентификации
		endpoints.GET("/get-profile", authMiddleware, h.GetProfile)
	}

	return authMiddleware
}
