package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/qreaqtor/api-avito-shop/internal/app"
	"github.com/qreaqtor/api-avito-shop/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	app, err := app.StartNewApp(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}
