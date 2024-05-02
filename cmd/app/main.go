// main.go

package main

import (
	"Cinema/internal/app"
	"Cinema/internal/config"
	"Cinema/pkg/common/logging"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logging.L(ctx).Info("config initializing")
	cfg := config.GetConfig()

	newApp, err := app.NewApp(ctx, cfg)
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.NewApp")
	}

	logging.L(ctx).Info("Running application")

	err = newApp.Run(ctx)
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.Run")
		return
	}

}
