package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/nala/v1")

	// Routes for GET method
	route.Get("/posts/:post_id", handlers.GetPost)
	route.Get("/posts/:lang", handlers.GetPostsByLanguage)
	route.Get("/posts/:usr_id", handlers.GetPostsByUser)
	route.Get("/posts/:cat", handlers.GetPostsByCategory)
	route.Get("/posts/:diff", handlers.GetPostsByDifficulty)
	route.Get("/posts/:media", handlers.GetPostsByMediaType)
	route.Get("/posts/tags/:id", handlers.GetPostTags)
	route.Get("/posts/likes/:id", handlers.GetPostLikes)

	// Routes for POST method
	route.Post("/posts", middleware.PasetoProtected(handlers.TokenManager), handlers.AddPost)
	route.Post("/posts/vote", middleware.PasetoProtected(handlers.TokenManager), handlers.VotePost)

	// Routes for PATCH method
	route.Patch("/posts", middleware.PasetoProtected(handlers.TokenManager), handlers.UpdatePost)

	// Routes for DELETE method
	route.Delete("/posts/:post_id", middleware.PasetoProtected(handlers.TokenManager), handlers.RemovePost)
	route.Delete("/posts/vote", middleware.PasetoProtected(handlers.TokenManager), handlers.RemoveVote)
}
