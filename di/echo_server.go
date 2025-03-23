package di

import (
	"context"
	// "os"
	// "os/signal"
	// "syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/labstack/gommon/log"
)

func setupServer(pctx context.Context, e *echo.Echo, c Config) {
	// Request Timeout
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Request Timeout",
		Timeout:      30 * time.Second,
	}))
}