package middleware

import (
	"fmt"
	"strings"

	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
)

func PasetoProtected(tokenManager validation.TokenManager) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get(util.AuthorizationHeaderKey)
		if len(authHeader) == 0 {
			return controllers.SendFailureResponse(
				ctx,
				fiber.StatusUnauthorized,
				"user is not authorized to access the content",
			)
		}

		authFields := strings.Fields(authHeader)

		fmt.Printf("Auth Header: %s; Auth Token: %s", authFields[0], authFields[1])

		if len(authFields) < 2 {
			return controllers.SendFailureResponse(
				ctx,
				fiber.StatusUnauthorized,
				"invalid authorization header format",
			)
		}

		authType := strings.ToLower(authFields[0])

		if authType != util.AuthorizationTypeBearer {
			return controllers.SendFailureResponse(
				ctx,
				fiber.StatusUnauthorized,
				"authorization type not supported",
			)
		}

		token := authFields[1]
		payload, err := tokenManager.VerifyToken(token)
		if err != nil {
			return controllers.SendFailureResponse(
				ctx,
				fiber.StatusForbidden,
				"token is invalid",
			)
		}
		ctx.Locals(util.AuthPayloadKey, payload)

		return ctx.Next()
	}
}
