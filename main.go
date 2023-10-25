package main

import (
	"flag"
	"hyphen-hellog/client/siss"
	"hyphen-hellog/database"
	"hyphen-hellog/ent"
	"hyphen-hellog/middleware"
	"hyphen-hellog/model/request"
	"hyphen-hellog/model/response"
	"hyphen-hellog/verifier"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var port string

func init() {
	flagPort := flag.Int("p", 8080, "Enter the port")
	flag.Parse()

	port = ":" + strconv.Itoa(*flagPort)
}

func main() {
	app := fiber.New(fiber.Config{})
	// app.Use(recover.New())
	app.Use(middleware.Auth)

	api := app.Group("/api/hellog")

	// 완
	api.Post("/post", func(c *fiber.Ctx) (err error) {
		clientRequest := new(request.CreatePost).Parse(c)
		verifier.Validate(c)

		database.New().CreatePostX(c.Context(),
			&ent.Post{
				Title:        clientRequest.Title,
				Content:      clientRequest.Content,
				PreviewImage: siss.CreateImage(clientRequest.PreviewImage),
				IsPrivate:    clientRequest.IsPrivate,
			},
			c.Locals("user").(*ent.Author).ID)

		return c.Status(fiber.StatusOK).JSON(response.Genreal{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data:    nil,
		})
	})

	api.Post("/:post_id/comment", func(c *fiber.Ctx) error {
		clientRequest := new(request.CreateComment).Parse(c)
		verifier.Validate(c)

		database.New().CreateComment(c.Context(),
			&ent.Comment{
				Content: clientRequest.Content,
			},
			clientRequest.ParentID,
			clientRequest.PostID,
			c.Locals("user").(*ent.Author).ID)

		return c.Status(fiber.StatusOK).JSON(response.Genreal{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data:    nil,
		})
	})

	api.Get("/:post_id", func(c *fiber.Ctx) error {
		clientRequest := new(request.GetPost).Parse(c)
		verifier.Validate(c)

		post := database.New().GetPostX(c.Context(), clientRequest.PostID)
		author := database.New().GetAuthorXByPostID(c.Context(), post.ID)

		return c.Status(fiber.StatusOK).JSON(response.Genreal{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data: response.GetPost{
				Post:   post,
				Author: author,
			},
		})
	})

	api.Get("/:post_id/comments", func(c *fiber.Ctx) error {
		return nil
	})

	api.Patch("/:post_id", func(c *fiber.Ctx) error {
		return nil
	})
	api.Patch("/:post_id/comment", func(c *fiber.Ctx) error {
		return nil
	})
	api.Patch("/:post_id/unliked", func(c *fiber.Ctx) error {
		return nil
	})
	api.Patch("/:post_id/liked", func(c *fiber.Ctx) error {
		return nil
	})

	api.Delete("/:post_id", func(c *fiber.Ctx) error {
		return nil
	})
	api.Delete("/:post_id/comment", func(c *fiber.Ctx) error {
		return nil
	})

	log.Fatal(app.Listen(port))

}
