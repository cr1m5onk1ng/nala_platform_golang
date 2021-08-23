package validation

import (
	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func ValidateVoteData(ctx *fiber.Ctx, vote *db.Vote) (*db.Vote, error) {
	if err := ctx.BodyParser(vote); err != nil {
		// Return status 400 and error message.
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return vote, nil
}

func ValidateVoteAndAuthorization(ctx *fiber.Ctx, vote *db.Vote) (*db.Vote, error) {
	if err := CheckUserAuthorization(ctx); err != nil {
		return nil, err
	}
	v, err := ValidateVoteData(ctx, vote)
	if err != nil {
		return nil, err
	}
	return v, nil
}
