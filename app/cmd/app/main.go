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
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg, &logger)
	if err != nil {
		log.Fatal(err)
	}

	a.Run()

}
