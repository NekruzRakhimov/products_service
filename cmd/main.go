package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/NekruzRakhimov/product_service/internal/bootstrap"
	"github.com/NekruzRakhimov/product_service/internal/config"
	"github.com/sethvargo/go-envconfig"
)

// @title ProductService API
// @contact.name ProductService API Service
// @contact.url http://test.com
// @contact.email test@test.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	var cfg config.Config

	err := envconfig.ProcessWith(context.TODO(), &envconfig.Config{Target: &cfg, Lookuper: envconfig.OsLookuper()})
	if err != nil {
		panic(err)
	}

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt)

	app := bootstrap.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quitSignal
		cancel()
	}()

	app.Run(ctx)

}
