package controllers

import (
	"os"
	"strconv"

	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
)

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]
func GetNewJwtAccessToken(c *fiber.Ctx) error {
	// Generate a new Access token.
	secret := os.Getenv("JWT_SECRET_KEY")
	minutesCount, err := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	if err != nil {
		return err
	}
	token, err := validation.GenerateNewAccessToken(secret, minutesCount)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
