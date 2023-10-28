package handlers

import (
	"hyphen-hellog/cerrors"
	"hyphen-hellog/client/siss"
	"hyphen-hellog/database"
	"hyphen-hellog/ent"
	"hyphen-hellog/model"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	clientRequest := new(model.InCreatePost).ParseX(c)

	database.Get().CreatePostX(c.Context(),
		&ent.Post{
			Title:        clientRequest.Title,
			Content:      clientRequest.Content,
			PreviewImage: siss.CreateImage(clientRequest.PreviewImage),
			IsPrivate:    clientRequest.IsPrivate,
		},
		c.Locals("user").(*ent.Author).ID)

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func GetPost(c *fiber.Ctx) error {
	clientRequest := new(model.InGetPost).ParseX(c)

	post := database.Get().GetPostX(c.Context(), clientRequest.PostID)

	author := database.Get().GetAuthorByPostIDX(c.Context(), post.ID)

	var isLiked bool
	if c.Locals("user").(*ent.Author) != nil {
		isLiked = database.Get().IsLikedByAuthorIDX(c.Context(), author.ID)
	} else {
		isLiked = false
	}

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data: model.OutGetPost{
			Post:    post,
			IsLiked: isLiked,
			Author:  author,
		},
	})
}

func UpdatePost(c *fiber.Ctx) error {
	clientRequest := new(model.InUpdatePost).ParseX(c)

	if database.Get().GetAuthorByPostIDX(c.Context(), clientRequest.PostID).ID != c.Locals("user").(*ent.Author).ID {
		panic(cerrors.WrongApproachErr{
			Err: "해당 사용자는 이 포스트의 주인이 아니므로 수정 하지 못합니다.",
		})
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

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func DeletePost(c *fiber.Ctx) error {
	clientRequest := new(model.InDeletePost).ParseX(c)

	if database.Get().GetAuthorByPostIDX(c.Context(), clientRequest.PostID).ID != c.Locals("user").(*ent.Author).ID {
		panic(cerrors.WrongApproachErr{
			Err: "해당 사용자는 이 포스트의 주인이 아니므로 삭제 하지 못합니다.",
		})
	}

	database.Get().DeletePostX(c.Context(), clientRequest.PostID)

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}
