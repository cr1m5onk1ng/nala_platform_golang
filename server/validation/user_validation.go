package validation

import (
	"time"

	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CheckUserDataValidty(ctx *fiber.Ctx, user *domain.MappedUser) (*domain.MappedUser, error) {
	// Check if data is valid
	if err := ctx.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	user.ID = uuid.NewString()
	/*
		// Create a new validator for a User model.
		validate := NewValidator()

		// Validate user fields.
		if err := validate.Struct(user); err != nil {
			// Return, if some fields are not valid.
			return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": ValidatorErrors(err),
				"data":    nil,
			})
		}

	*/
	return user, nil
}

func CheckUserAuthorization(ctx *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	expires := claims.Expires

	if now > expires {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "unauthorized, token expired",
			"data":    nil,
		})
	}

	return nil
}

func CheckUserDataValidtyAndAuthorization(ctx *fiber.Ctx, user *domain.MappedUser) (*domain.MappedUser, error) {
	var err error
	var userToCheck *domain.MappedUser

	err = CheckUserAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	userToCheck, err = CheckUserDataValidty(ctx, user)
	if err != nil {
		return nil, err
	}
	return userToCheck, nil
}
