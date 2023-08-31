package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"reqwizard/pkg/postgres"
	"reqwizard/pkg/postgres/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"reqwizard/internal/routes"
	service_email "reqwizard/internal/services/email"
)

type App struct {
	httpServer *http.Server

	pgGorm *gorm.Gorm
	mailer *service_email.Mailer
}

func New(ctx context.Context) (*App, error) {
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
		pgGorm: pgGorm,
		mailer: mailer,
	}, nil
}

func (app *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("Content-Type")
	router.Use(
		cors.New(corsConfig),
		gin.Recovery(),
		gin.Logger(),
	)

	routes.InitRoutes(router, app.pgGorm, app.mailer)

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
