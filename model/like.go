package model

import (
	"hyphen-hellog/cerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type InUpdateLike struct {
	PostID int
}

func (i *InUpdateLike) ParseX(c *fiber.Ctx) *InUpdateLike {
	var err error

	i.PostID, err = strconv.Atoi(c.Params("post_id"))
	if err != nil {
		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	return i
}

type InUpdateUnlike struct {
	PostID int
}

func (i *InUpdateUnlike) ParseX(c *fiber.Ctx) *InUpdateUnlike {
	var err error

	i.PostID, err = strconv.Atoi(c.Params("post_id"))
	if err != nil {
		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	return i
}
