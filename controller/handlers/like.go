package handlers

import (
	"hyphen-hellog/database"
	"hyphen-hellog/ent"
	"hyphen-hellog/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateUnliked(c *fiber.Ctx) error {
	postID, err := strconv.Atoi(c.Params("post_id"))
	if err != nil {
		panic(err)
	}

	database.Get().DeleteLikeX(c.Context(), c.Locals("user").(*ent.Author).ID, postID)

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func UpdateLiked(c *fiber.Ctx) error {
	postID, err := strconv.Atoi(c.Params("post_id"))
	if err != nil {
		panic(err)
	}

	authorID := c.Locals("user").(*ent.Author).ID
	_, err = database.Get().UpdateLike(c.Context(), authorID, postID)
	if err != nil {
		database.Get().CreateLikeX(c.Context(), authorID, postID)
	}

	return c.Status(fiber.StatusOK).JSON(response.Genreal{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}
