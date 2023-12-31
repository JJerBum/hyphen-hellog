package handlers

import (
	"fmt"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/client/siss"
	"hyphen-hellog/database"
	"hyphen-hellog/ent"
	"hyphen-hellog/model"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	clientRequest := new(model.InCreatePost).ParseX(c)

	var previewImage string = ""
	if clientRequest.PreviewImage != nil {
		previewImage = siss.CreateImage(clientRequest.PreviewImage)
	}

	database.Get().CreatePostX(c.Context(),
		&ent.Post{
			Title:            clientRequest.Title,
			Content:          clientRequest.Content,
			ShortDescription: clientRequest.ShortDescription,
			PreviewImage:     previewImage,
			IsPrivate:        clientRequest.IsPrivate,
		},
		c.Locals("user").(*ent.Author).ID)

	return c.Status(fiber.StatusCreated).JSON(model.General{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    nil,
	})
}

func GetPosts(c *fiber.Ctx) error {

	response := new(model.OutGetPosts)
	response.Posts = make([]*model.OutGetPost, 0)

	posts := database.Get().GetPostsX(c.Context())

	for _, post := range posts {
		author, err := database.Get().GetAuthorByPostID(c.Context(), post.ID)
		if err != nil {
			panic(cerrors.WrongApproachErr{Err: fmt.Sprintf("postID가 %d인 post를 찾을 수 없습니다.", post.ID)})
		}

		var isLiked bool
		if c.Locals("user").(*ent.Author) != nil {
			isLiked = database.Get().IsLikedByAuthorIDX(c.Context(), author.AuthorID)
		} else {
			isLiked = false
		}

		response.Posts = append(response.Posts, &model.OutGetPost{
			Post:    post,
			Author:  author,
			IsLiked: isLiked,
			MyLikes: database.Get().GetPostMyLikesX(c.Context(), post.ID),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
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
			MyLikes: database.Get().GetPostMyLikesX(c.Context(), post.ID),
		},
	})
}

func UpdatePost(c *fiber.Ctx) error {
	clientRequest := new(model.InUpdatePost).ParseX(c)

	if database.Get().GetAuthorByPostIDX(c.Context(), clientRequest.PostID).ID != c.Locals("user").(*ent.Author).ID {
		panic(cerrors.UnauthorizedErr{
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
		panic(cerrors.UnauthorizedErr{
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
