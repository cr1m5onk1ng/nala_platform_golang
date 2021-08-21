package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/api/v1")

	// Routes for GET method:
	route.Get("/users", handlers.GetAllUsers)
	route.Get("/users/:id", handlers.GetUser)

	// Routes for POST method:
	//route.Post("/users", middleware.JWTProtected(), handlers.CreateUser)
	route.Post("/users", handlers.CreateUserNotSecure)

	// Routes for PATCH method
	route.Patch("/users/:id", middleware.JWTProtected(), handlers.GetUser)

}
