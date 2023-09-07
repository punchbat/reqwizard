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
	"github.com/spf13/viper"
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
	clientsMap := viper.GetStringMap("clients")
	clientOrigins := []string{}
	web, ok := clientsMap["web"]
	if ok {
		clientOrigins = append(clientOrigins, web.(string))
	}

	// * CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = viper.GetBool("cors.credentials")
	corsHeaders := viper.GetStringSlice("cors.headers")
	corsConfig.AllowAllOrigins = !(len(corsHeaders) > 0)
	corsConfig.AllowOrigins = clientOrigins
	corsConfig.AddAllowHeaders(corsHeaders...)
	// * ALLOW TO GET FROM BROWSER HEADER
	corsConfig.AddExposeHeaders("X-Csrf-Token")

	// * CSRF
	csrfMiddleware := csrf.Protect(
		[]byte(viper.GetString("csrf.auth_key")),
		csrf.Secure(viper.GetBool("csrf.secure")),
		csrf.TrustedOrigins(clientOrigins),
		csrf.Path(viper.GetString("csrf.path")),
		csrf.HttpOnly(viper.GetBool("csrf.http_only")),
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
