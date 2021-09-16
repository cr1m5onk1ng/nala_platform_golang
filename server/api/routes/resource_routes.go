package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ResourceRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/nala/v1")

	// Routes for GET method:
	route.Get("/resources/:id", handlers.GetResource)
	route.Get("/resources/langs/:lang", handlers.GetResourcesByLanguage)
	route.Get("/resources/urls/:url", handlers.GetResourceByUrl)

	// Routes for POST method:
	//route.Post("/resources", middleware.JWTProtected(), handlers.AddResource)
	route.Post("/resources", middleware.PasetoProtected(handlers.TokenManager), handlers.AddResource)

}
