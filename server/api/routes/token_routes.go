package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func TokenRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/nala/v1")

	route.Post("/users/tokens/jwt", middleware.JWTProtected(), handlers.GetNewJwtAccessToken)

	route.Post("/users/tokens/paseto", middleware.PasetoProtected(handlers.TokenManager), handlers.GetNewPasetoAccessToken)

}
