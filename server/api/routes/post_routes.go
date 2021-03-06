package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/nala/v1")

	// Routes for GET method
	route.Get("/posts/q", handlers.GetPosts)
	route.Get("/posts/:id", handlers.GetPost)
	route.Get("/posts/langs/:lang", handlers.GetPostsByLanguage)
	route.Get("/posts/users/:user", handlers.GetPostsByUser)
	route.Get("/posts/categories/:cat", handlers.GetPostsByCategory)
	route.Get("/posts/difficulties/:diff", handlers.GetPostsByDifficulty)
	route.Get("/posts/medias/:media", handlers.GetPostsByMediaType)
	route.Get("/posts/:post/tags", handlers.GetPostTags)
	route.Get("/posts/:post/likes", handlers.GetPostLikes)

	// Routes for POST method
	route.Post("/posts", middleware.PasetoProtected(handlers.TokenManager), handlers.AddPost)
	route.Post("/posts/:post/vote", middleware.PasetoProtected(handlers.TokenManager), handlers.VotePost)

	// Routes for PUT method
	route.Put("/posts", middleware.PasetoProtected(handlers.TokenManager), handlers.UpdatePost)

	// Routes for DELETE method
	route.Delete("/posts/:post", middleware.PasetoProtected(handlers.TokenManager), handlers.RemovePost)
	route.Delete("/posts/:post/vote", middleware.PasetoProtected(handlers.TokenManager), handlers.RemoveVote)
}
