package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"reqwizard/pkg/postgres"
	"reqwizard/pkg/postgres/gorm"

	"reqwizard/internal/routes"
	service_email "reqwizard/internal/services/email"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/csrf"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type App struct {
	httpServer *http.Server

	c      *cron.Cron
	pgGorm *gorm.Gorm
	mailer *service_email.Mailer
}

func New(ctx context.Context) (*App, error) {
	c := cron.New()

	err := postgres.RunMigrations()
	if err != nil {
		return nil, err
	}

	pgGorm, err := gorm.New(ctx)
	if err != nil {
		return nil, err
	}

	mailer := service_email.NewMailer()

	return &App{
		c:      c,
		pgGorm: pgGorm,
		mailer: mailer,
	}, nil
}

func (app *App) Run(port string) error {
	go func() {
		app.c.Start()
	}()

	router := gin.Default()

	// * CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	// corsConfig.AllowAllOrigins = true
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowOrigins = []string{"http://localhost:8000"}
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("Content-Type")


	// * CSRF 
	csrfMiddleware := csrf.Default(csrf.Options{
        Secret: "your-secret-key", // Замените на ваш секретный ключ
    })

	router.Use(
		cors.New(corsConfig),
		csrfMiddleware,
		gin.Recovery(),
		gin.Logger(),
	)

	// * ROUTES
	routes.InitRoutes(router, app.c, app.pgGorm, app.mailer)

	// Конфиги для сервера
	app.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("App started, listen port: %s", port)

	if err := app.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (app *App) Stop() error {
	app.c.Stop()
	<-app.c.Stop().Done()
	log.Println("Cron jobs gracefully stopped")

	err := app.pgGorm.DB.Close()
	if err == nil {
		log.Println("pg gorm connection gracefully stopped")
	} else {
		log.Fatalf("failed to close pg gorm connection. %s", err.Error())
	}

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
