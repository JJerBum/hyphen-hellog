package handlers

import (
	"hyphen-hellog/database"
	"hyphen-hellog/ent"
	"hyphen-hellog/model"

	"github.com/gofiber/fiber/v2"
)

func UpdateUnliked(c *fiber.Ctx) error {
	clientReqeust := new(model.InUpdateUnlike).ParseX(c)

	database.Get().DeleteLikeX(c.Context(), c.Locals("user").(*ent.Author).ID, clientReqeust.PostID)

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func UpdateLiked(c *fiber.Ctx) (err error) {
	clientReqeust := new(model.InUpdateLike).ParseX(c)

	authorID := c.Locals("user").(*ent.Author).ID
	_, err = database.Get().UpdateLike(c.Context(), authorID, clientReqeust.PostID)
	if err != nil {
		database.Get().CreateLikeX(c.Context(), authorID, clientReqeust.PostID)
	}

	return c.Status(fiber.StatusOK).JSON(model.General{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})

}
