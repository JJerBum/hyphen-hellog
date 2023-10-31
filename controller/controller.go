package controller

import (
	"hyphen-hellog/controller/handlers"
	"hyphen-hellog/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) *fiber.App {
	api := app.Group("/api/hellog")

	// post
	api.Post("/posts/:post_id", middleware.RequireAuth, handlers.CreatePost)
	api.Get("/posts", middleware.Auth, handlers.GetPosts)
	api.Get("/posts/:post_id", middleware.Auth, handlers.GetPost)
	api.Patch("/posts/:post_id", middleware.RequireAuth, handlers.UpdatePost)
	api.Delete("/posts/:post_id", middleware.RequireAuth, handlers.DeletePost)

	// comment
	api.Post("/posts/:post_id/comments/comment", middleware.RequireAuth, handlers.CreateComment)
	api.Get("posts/:post_id/comments", middleware.RequireAuth, handlers.GetComment)
	api.Patch("/posts/comments/comment", middleware.RequireAuth, handlers.UpdateComment)
	api.Delete("/posts/comments/comment", middleware.RequireAuth, handlers.DeleteComment)

	// like
	api.Patch("/posts/:post_id/unlike", middleware.RequireAuth, handlers.UpdateUnliked)
	api.Patch("/posts/:post_id/like", middleware.RequireAuth, handlers.UpdateLiked)

	return app
}
