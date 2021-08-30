package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/nala/v1")

	// Routes for GET method:
	route.Get("/users", middleware.PasetoProtected(handlers.TokenManager), handlers.GetAllUsers)
	route.Get("/users/:id", middleware.PasetoProtected(handlers.TokenManager), handlers.GetUser)

	// Routes for POST method:
	route.Post("/users", handlers.CreateUser)

	route.Post("/users/login", handlers.LoginUser)

	route.Post("/users/targetlang", middleware.PasetoProtected(handlers.TokenManager), handlers.AddUserTargetLanguage)

	// Routes for PATCH method
	route.Patch("/users/:id/lang", middleware.PasetoProtected(handlers.TokenManager), handlers.UpdateUserLanguage)

	route.Patch("/users/:id/role", middleware.PasetoProtected(handlers.TokenManager), handlers.UpdateUserRole)

}
