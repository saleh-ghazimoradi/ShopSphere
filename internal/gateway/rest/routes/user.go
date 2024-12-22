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

func UserRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	authService := helper.NewAuth(config.AppConfig.AppSecret.Secret)
	userService := service.NewUserService(userRepository, authService)
	user := handlers.NewUserHandler(userService, authService)

	// Public routes
	pubRoutes := app.Group("/users")
	pubRoutes.Post("/register", user.Register)
	pubRoutes.Post("/login", user.Login)

	// Private routes
	pvtRoutes := pubRoutes.Group("/", authService.Authorize)
	pvtRoutes.Get("/verify", user.GetVerificationCode)
	pvtRoutes.Post("/verify", user.Verify)
	pvtRoutes.Post("/profile", user.CreateProfile)
	pvtRoutes.Get("/profile", user.GetProfile)
	pvtRoutes.Post("/cart", user.AddToCart)
	pvtRoutes.Get("/cart", user.GetCart)
	pvtRoutes.Get("/order", user.GetOrders)
	pvtRoutes.Get("/order/:id", user.GetOrder)
	pvtRoutes.Post("/become-seller", user.BecomeSeller)
}
