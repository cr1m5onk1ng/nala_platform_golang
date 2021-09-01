package routes

import (
	"github.com/cr1m5onk1ng/nala_platform_app/api/controllers"
	"github.com/cr1m5onk1ng/nala_platform_app/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func DiscussionsRoutes(app *fiber.App, handlers *controllers.Handlers) {
	route := app.Group("/nala/v1")

	// Routes for GET method
	route.Get("/posts/discussions/:post", handlers.GetPostDiscussions)
	route.Get("/posts/discussions/comments/:discussion", handlers.GetDiscussionComments)
	route.Get("/posts/discussions/comments/likes/:comment", handlers.GetCommentLikes)
	route.Get("/posts/discussions/comments/likes/count/:comment", handlers.GetCommentLikesCount)

	// Routes for POST method
	route.Post("/posts/discussions", middleware.PasetoProtected(handlers.TokenManager), handlers.CreateDiscussion)
	route.Post("/posts/discussions/comments", middleware.PasetoProtected(handlers.TokenManager), handlers.AddCommentOnDiscussion)
	route.Post("/posts/discussions/comments/like", middleware.PasetoProtected(handlers.TokenManager), handlers.LikeComment)

	// Routes for PUT method
	route.Put("/posts/discussions", middleware.PasetoProtected(handlers.TokenManager), handlers.UpdateDiscussion)
	route.Put("/posts/discussions/comments", middleware.PasetoProtected(handlers.TokenManager), handlers.UpdateDiscussionComment)

	// Routes for DELETE method
	route.Delete("/posts/discussions", middleware.PasetoProtected(handlers.TokenManager), handlers.RemoveDiscussion)
	route.Delete("/posts/discussions/comments", middleware.PasetoProtected(handlers.TokenManager), handlers.RemoveCommentFromDiscussion)
	route.Delete("/posts/discussions/comments/like", middleware.PasetoProtected(handlers.TokenManager), handlers.UnlikeComment)
}
