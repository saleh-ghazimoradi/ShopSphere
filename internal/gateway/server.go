package gateway

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ShopSphere/config"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/gateway/rest/routes"
	"github.com/saleh-ghazimoradi/ShopSphere/utils"
)

func Server() error {
	app := fiber.New()

	db, err := utils.DBConnection(utils.DBMigrator)
	if err != nil {
		return err
	}

	routes.HealthCheck(app)
	routes.UserRoutes(app, db)

	if err = app.Listen(config.AppConfig.ServerConfig.Port); err != nil {
		return err
	}
	return nil
}
