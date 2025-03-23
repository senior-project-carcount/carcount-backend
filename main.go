// Package main is the main entry point for shiftover-backend service
package main

import (
	"os"

	"github.com/labstack/gommon/log"

	"github.com/senior-project-carcount/carcount-backend/config"
	"github.com/senior-project-carcount/carcount-backend/di"
)

func main() {
	env := os.Getenv("APP_ENV_STAGE")
	if env == "" {
		env = "dev"
	}

	// Initiaize Config
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Panicf("error - [main.loadConfig] error loading config: %v", err)
	}

	di.New(di.Config{
		AppConfig:   cfg.AppConfig,
		MongoConfig: cfg.MongoConfig,
	})
}
