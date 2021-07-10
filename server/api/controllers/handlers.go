package controllers

import (
	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Repo *db.Repository
}

func NewHandlers(repo *db.Repository) *Handlers {
	return &Handlers{Repo: repo}
}

func SendSuccessResponse(ctx *fiber.Ctx, statusCode int, payload interface{}) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    payload,
	})
}

func SendFailureResponse(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"error":   true,
		"message": message,
		"data":    nil,
	})
}
