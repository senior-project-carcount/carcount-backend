// Package di provides dependency injection for the server
package di

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/senior-project-carcount/carcount-backend/config"
	"github.com/senior-project-carcount/carcount-backend/handler"
)

// Config represents the configuration of the service
type Config struct {
	AppConfig   config.AppConfig
	MongoConfig config.MongoConfig
}

// New injects the dependencies for the server
func New(c Config) {
	ctx := context.Background()

	e := echo.New()
	setupServer(ctx, e, c)
	// mongoRepos := newMongoRepositories(ctx, c)

	handler.New(e, handler.Dependencies{
		// Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.initServer] unable to start server: %v", err)
	}
}
