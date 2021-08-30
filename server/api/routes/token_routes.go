package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func TokenRoutes(app *fiber.App) {
	route := app.Group("/nala/v1")

	route.Get("/token/jwt", controllers.GetNewJwtAccessToken)

}
