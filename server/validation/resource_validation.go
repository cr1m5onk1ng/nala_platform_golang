package validation

import (
	"fmt"
	"net/url"

	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/gofiber/fiber/v2"
)

func ValidateResourceUrl(urlToParse string) error {
	_, err := url.ParseRequestURI(urlToParse)
	if err != nil {
		return fmt.Errorf("error while parsing url: %s", urlToParse)
	}
	return nil
}

func ValidateResourceData(ctx *fiber.Ctx, resource *domain.MappedResource) (*domain.MappedResource, error) {
	// Check if data is valid
	if err := ctx.BodyParser(resource); err != nil {
		// Return status 400 and error message.
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	// Create a new validator for a User model.
	validate := NewValidator()

	// Validate user fields.
	if err := validate.Struct(resource); err != nil {
		// Return, if some fields are not valid.
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": ValidatorErrors(err),
			"data":    nil,
		})
	}
	return resource, nil
}

func ValidateResourceDataAndUrlAndAuthorization(ctx *fiber.Ctx, resource *domain.MappedResource) (*domain.MappedResource, error) {
	var err error
	if err = CheckUserAuthorization(ctx); err != nil {
		return nil, err
	}

	r, err := ValidateResourceData(ctx, resource)
	if err != nil {
		return nil, err
	}

	if err = ValidateResourceUrl(r.Url); err != nil {
		return nil, err
	}

	return resource, nil
}
