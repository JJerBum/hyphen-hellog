package handlers

import (
	"hyphen-hellog/client/siss"
	"hyphen-hellog/database"
	"hyphen-hellog/ent"
	"hyphen-hellog/model/request"
	"hyphen-hellog/model/response"
	"hyphen-hellog/verifier"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	clientRequest := new(request.CreatePost).Parse(c)
	verifier.Validate(c)

	database.Get().CreatePostX(c.Context(),
		&ent.Post{
			Title:        clientRequest.Title,
			Content:      clientRequest.Content,
			PreviewImage: siss.CreateImage(clientRequest.PreviewImage),
			IsPrivate:    clientRequest.IsPrivate,
		},
		c.Locals("user").(*ent.Author).ID)

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func GetPost(c *fiber.Ctx) error {
	clientRequest := new(request.GetPost).Parse(c)
	verifier.Validate(c)

	post := database.Get().GetPostX(c.Context(), clientRequest.PostID)

	author := database.Get().GetAuthorXByPostID(c.Context(), post.ID)

	var isLiked bool
	if c.Locals("user").(*ent.Author) != nil {
		isLiked = database.Get().IsLikedXByAuthorID(c.Context(), author.ID)
	} else {
		isLiked = false
	}

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data: response.GetPost{
			Post:    post,
			IsLiked: isLiked,
			Author:  author,
		},
	})
}

func UpdatePost(c *fiber.Ctx) error {
	clientRequest := new(request.UpdatePost).Parse(c)
	verifier.Validate(c)

	if database.Get().GetAuthorXByPostID(c.Context(), clientRequest.PostID).ID != c.Locals("user").(*ent.Author).ID {
		panic("잘못된 접근 입니다. (해당 사용자는 해당 포스트를 변경할 수 없습니다.)")
	}

	// 왼래 있었던 이미지 삭제
	siss.DeleteImage(database.Get().GetPostX(c.Context(), clientRequest.PostID).PreviewImage)

	// 업데이트
	database.Get().UpdatePostX(c.Context(),
		&ent.Post{
			ID:           clientRequest.PostID,
			Title:        clientRequest.Title,
			Content:      clientRequest.Content,
			PreviewImage: siss.CreateImage(clientRequest.PreviewImage),
			IsPrivate:    clientRequest.IsPrivate,
		},
		c.Locals("user").(*ent.Author).ID)

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func DeletePost(c *fiber.Ctx) error {
	clientRequest := new(request.DeletePost).Parse(c)

	if database.Get().GetAuthorXByPostID(c.Context(), clientRequest.PostID).ID != c.Locals("user").(*ent.Author).ID {
		panic("잘못된 접근 입니다. (해당 사용자는 해당 포스트를 삭제할 수 없습니다.)")
	}

	database.Get().DeletePostX(c.Context(), clientRequest.PostID)

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}
