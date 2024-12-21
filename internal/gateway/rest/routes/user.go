package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/gateway/rest/handlers"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/repository"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service"
	"gorm.io/gorm"
)

func UserRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	user := handlers.NewUserHandler(userService)

	app.Post("/register", user.Register)
	app.Post("/login", user.Login)
	app.Get("/verify", user.GetVerificationCode)
	app.Post("/verify", user.Verify)
	app.Post("/profile", user.CreateProfile)
	app.Get("/profile", user.GetProfile)
	app.Post("/cart", user.AddToCart)
	app.Get("/cart", user.GetCart)
	app.Get("/order", user.GetOrders)
	app.Get("/order/:id", user.GetOrder)
	app.Post("/become-seller", user.BecomeSeller)
}
