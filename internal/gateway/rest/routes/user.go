package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/gateway/rest/handlers"
)

func UserRoutes(app *fiber.App) {
	user := handlers.NewUserHandler()
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
