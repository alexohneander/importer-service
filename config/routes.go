package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func ConfigureRoutes(app *fiber.App) {

	// Monitoring Dashboard
	app.Get("/dashboard", monitor.New())
}
