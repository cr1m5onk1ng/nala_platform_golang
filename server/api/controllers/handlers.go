package controllers

import (
	rp "github.com/cr1m5onk1ng/nala_platform_app/repository"
	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Repo         rp.Repository
	TokenManager validation.TokenManager
	Config       util.Config
}

func NewHandlers(config util.Config, repo rp.Repository, tokenManager validation.TokenManager) (*Handlers, error) {
	return &Handlers{Config: config, Repo: repo, TokenManager: tokenManager}, nil
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
