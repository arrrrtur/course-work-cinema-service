package main

import (
	"Cinema/internal/app"
	"Cinema/internal/config"
	"Cinema/pkg/logging"
	"log"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)

	logger := logging.GetLogger()

	a, err := app.NewApp(cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	a.Run()

}
