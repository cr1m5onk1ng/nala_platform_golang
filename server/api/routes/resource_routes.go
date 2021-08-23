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
	route.Get("/resource/:lang", handlers.GetResourcesByLanguage)
	route.Get("/resource/:url", handlers.GetResourceByUrl)
	route.Get("/resource/:postId", handlers.GetResourcePost)
	route.Get("/resource/:usrId", handlers.GetResourcesPostsByUser)

	// Routes for POST method:
	//route.Post("/resources", middleware.JWTProtected(), handlers.AddResource)
	route.Post("/resources", handlers.AddResourceNotSecure)

	// Routes for PATCH method
	route.Patch("/resources/:id", middleware.JWTProtected(), handlers.UpdateResource)

}
