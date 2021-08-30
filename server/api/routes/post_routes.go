package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/nala/v1")

	// Routes for GET method
	route.Get("/posts/:post-id", handlers.GetPost)
	route.Get("/posts/:lang", handlers.GetPostsByLanguage)
	route.Get("/posts/:usr-id", handlers.GetPostsByUser)
	route.Get("/posts/:cat", handlers.GetPostsByCategory)
	route.Get("/posts/:diff", handlers.GetPostsByDifficulty)
	route.Get("/posts/:media", handlers.GetPostsByMediaType)
	route.Get("/posts/post-id/tags", handlers.GetPostTags)
	route.Get("/posts/post-id/likes", handlers.GetPostLikes)

	// Routes for POST method
	route.Post("/posts", middleware.PasetoProtected(handlers.TokenManager), handlers.AddPost)
	route.Post("/posts/:post-id/vote", middleware.PasetoProtected(handlers.TokenManager), handlers.VotePost)

	// Routes for PATCH method
	route.Put("/posts/:post-id", middleware.PasetoProtected(handlers.TokenManager), handlers.UpdatePost)

	// Routes for DELETE method
	route.Delete("/posts/:post-id", middleware.PasetoProtected(handlers.TokenManager), handlers.RemovePost)
	route.Delete("/posts/:post-id/vote", middleware.PasetoProtected(handlers.TokenManager), handlers.RemoveVote)
}
