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
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
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

	// * CLIENT ORIGINS
	clientOrigins := []string{"http://localhost:8000"}

	// * CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowOrigins = clientOrigins
	corsConfig.AddAllowHeaders("Authorization", "")
	corsConfig.AddAllowHeaders("Content-Type")
	corsConfig.AddAllowHeaders("X-CSRF-Token")
	// * ALLOW TO GET FROM BROWSER HEADER
	corsConfig.AddExposeHeaders("X-Csrf-Token")

	// * CSRF
	csrfMiddleware := csrf.Protect(
		[]byte("3d34b27f1df5c6ad7a24ac2fc7b0b340c5f1a88e27a22e44a73f129d3f0e9e6f"),
		csrf.Secure(false),
		csrf.TrustedOrigins(clientOrigins),
		csrf.Path("/"),
	)

	// * MIDDLEWARES
	router.Use(
		cors.New(corsConfig),
		gin.Recovery(),
		gin.Logger(),
	)

	// * CSRF
	router.Use(
		adapter.Wrap(csrfMiddleware),
	)

	router.Use(func(c *gin.Context) {
		c.Header("X-CSRF-Token", csrf.Token(c.Request))

		c.Next()
	})

	// * ROUTES
	routes.InitRoutes(router, app.c, app.pgGorm, app.mailer)

	// * Конфиги для сервера
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
