package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/gateway/rest/handlers"
)

func HealthCheck(app *fiber.App) {
	health := handlers.NewHealthHandler()
	app.Get("/health", health.Health)
}
