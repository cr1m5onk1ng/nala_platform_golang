package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/api/v1")

	// Routes for GET method
	route.Get("/posts/:postId", handlers.GetPost)
	route.Get("/posts/:lang", handlers.GetPostsByLanguage)
	route.Get("/posts/:usrId", handlers.GetPostsByUser)
	route.Get("/posts/:cat", handlers.GetPostsByCategory)
	route.Get("/posts/:diff", handlers.GetPostsByDifficulty)
	route.Get("/posts/:media", handlers.GetPostsByMediaType)
	route.Get("/posts/tags/:id", handlers.GetPostTags)
	route.Get("/posts/likes/:id", handlers.GetPostLikes)

	// Routes for POST method
	//route.Post("/posts", middleware.JWTProtected(), handlers.AddPost)
	route.Post("/posts", handlers.AddPostNotSecure)

	// Routes for PATCH method
	route.Patch("/posts/:postId", middleware.JWTProtected(), handlers.UpdatePost)

	// Routes for DELETE method
	route.Delete("/posts/:postId", middleware.JWTProtected(), handlers.RemovePost)
}
