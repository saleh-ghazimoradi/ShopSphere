package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/dto"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/helper"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service"
	"net/http"
)

type UserHandler struct {
	userService service.User
	authService helper.Auth
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {

	user := dto.UserSignUp{}

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "please provide valid input",
		})
	}

	token, err := u.userService.Signup(ctx.Context(), user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error signing up",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}

	if err := ctx.BodyParser(&loginInput); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "please provide valid input",
		})
	}

	token, err := u.userService.Login(ctx.Context(), loginInput.Email, loginInput.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "please provide correct user id password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token":   token,
	})

}

func (u *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func (u *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func (u *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func (u *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user, err := u.authService.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error get current user",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"user":    user,
	})
}

func (u *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func (u *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func (u *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{})
}

func (u *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func (u *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{})
}
func (u *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
	})
}

func NewUserHandler(userService service.User, authService helper.Auth) *UserHandler {
	return &UserHandler{
		userService: userService,
		authService: authService,
	}
}
