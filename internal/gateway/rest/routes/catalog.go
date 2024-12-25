package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ShopSphere/config"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/helper"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/repository"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service"
	"gorm.io/gorm"
)

func CatalogRoutes(app *fiber.App, db *gorm.DB) {
	catalogRepository := repository.NewCatalogRepository(db)
	catalogService := service.NewCatalogService(catalogRepository)
	authService := helper.NewAuth(config.AppConfig.AppSecret.Secret)
	catalog := handlers.NewCatalogHandler(catalogService, authService)

	// Listing Products and Categories
	app.Get("/products", catalog.GetProducts)
	app.Get("/products/:id", catalog.GetProduct)
	app.Get("/categories", catalog.GetCategories)
	app.Get("/categories/:id", catalog.GetCategoryById)

	sellerRoutes := app.Group("/seller", authService.AuthorizeSeller)
	// Categories
	sellerRoutes.Post("/categories", catalog.CreateCategories)
	sellerRoutes.Patch("/categories/:id", catalog.EditCategory)
	sellerRoutes.Delete("categories/:id", catalog.DeleteCategory)

	// Products
	sellerRoutes.Post("/products", catalog.CreateProducts)
	sellerRoutes.Get("/products", catalog.GetProducts)
	sellerRoutes.Get("/products/:id", catalog.GetProduct)
	sellerRoutes.Put("/products/:id", catalog.EditProduct)
	sellerRoutes.Patch("/products/:id", catalog.UpdateStock)
	sellerRoutes.Delete("/products/:id", catalog.DeleteProduct)
}
