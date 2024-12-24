package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func errorMessage(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(err.Error())
}

func internalError(ctx *fiber.Ctx, err error) error {
	return errorMessage(ctx, http.StatusInternalServerError, err)
}

func successMessage(ctx *fiber.Ctx, msg string, data any) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": msg,
		"data":    data,
	})
}
