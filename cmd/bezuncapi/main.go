package main

import (
	"bezuncapi/internal/app"
	"bezuncapi/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	// Echo instance
	e := echo.New()

	// Application configs
	configs := config.New(e.Logger)

	// Middleware
	app.Middleware(e)

	// Routes
	app.Routes(e.Router())

	if configs.Debug {
		e.Logger.SetLevel(log.DEBUG)
	}

	// Start server
	e.Logger.Fatal(e.Start(configs.ApplicationAddress()))
}
