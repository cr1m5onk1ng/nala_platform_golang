package validation

import (
	"github.com/cr1m5onk1ng/nala_platform_app/constants"
	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	validate := validator.New()

	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

func IsLanguageStringValid(lang string) bool {
	return util.CheckStringInSlice(constants.LANGUAGES(), lang)
}

func ValidateBodyData(ctx *fiber.Ctx, model interface{}) (interface{}, error) {
	if err := ctx.BodyParser(model); err != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return model, nil
}
