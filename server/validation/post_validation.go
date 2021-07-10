package validation

import (
	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ValidatePostData(ctx *fiber.Ctx, post *db.UserPost) (*db.UserPost, error) {
	// Check if data is valid
	if err := ctx.BodyParser(post); err != nil {
		// Return status 400 and error message.
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	post.ID = uuid.NewString()

	return post, nil
}

func ValidatePostDataAndAuthorization(ctx *fiber.Ctx, post *db.UserPost) (*db.UserPost, error) {
	var err error
	if err = CheckUserAuthorization(ctx); err != nil {
		return nil, err
	}

	p, err := ValidatePostData(ctx, post)
	if err != nil {
		return nil, err
	}

	return p, nil
}
