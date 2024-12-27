package gateway

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	c := cors.New(cors.Config{
		AllowOrigins: "http://localhost:3030",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	})

	app.Use(c)

	routes.HealthCheck(app)
	routes.UserRoutes(app, db)
	routes.CatalogRoutes(app, db)

	if err = app.Listen(config.AppConfig.ServerConfig.Port); err != nil {
		return err
	}
	return nil
}
