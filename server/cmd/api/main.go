package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"reqwizard/configs"
	"reqwizard/internal/app"
	"syscall"

	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)


func main() {
	// Загружаем конфиги, они доступны через viper.Get(...)
	configs.Init()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("%s", err.Error())
		os.Exit(1)
	}

	errGroup, errGroupCtx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		err := app.Run(viper.GetString("app.port"))
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		return err
	})

	select {
		case <-ctx.Done():
			log.Println("stop signal received")
			log.Println("app stopping...")
			app.Stop()
		case <-errGroupCtx.Done():
			log.Println("app stopping...")
			app.Stop()
	}

	log.Println("app stopped")
}