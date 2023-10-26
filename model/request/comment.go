package request

import (
	"hyphen-hellog/cerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreateComment 클라이언트로부터 요청된 바디 값과 매핑되는 구조체 입니다.
// PostID는 Params() 함수를 이용해야 합니다해
// PostID와 ParentID는 BodyParser()함수를 이용해야 합니다.
type CreateComment struct {
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	ParentID int    `json:"parent_id"`
}

func (c *CreateComment) Parse(ctx *fiber.Ctx) *CreateComment {
	var err error

	err = ctx.BodyParser(c)
	c.PostID, err = strconv.Atoi(ctx.Params("post_id"))

	if err != nil {
		panic(cerrors.ErrInvalidRequest)
	}

	return c
}

// GetComments 클라이언트로부터 요청된 바디 값과 매핑되는 구조체 입니다.
// PostID는 Params() 함수를 이용해야 합니다.
type GetComments struct {
	PostID int `json:"post_id"`
}

func (g *GetComments) Parse(ctx *fiber.Ctx) *GetComments {
	var err error

	g.PostID, err = strconv.Atoi(ctx.Params("post_id"))

	if err != nil {
		panic(cerrors.ErrInvalidRequest)
	}

	return g
}

type UpdateComment struct {
	CommentID int    `json:"comment_id"`
	Content   string `json:"content"`
}

func (u *UpdateComment) Parse(ctx *fiber.Ctx) *UpdateComment {
	var err error

	err = ctx.BodyParser(u)

	if err != nil {
		panic(cerrors.ErrInvalidRequest)
	}

	return u
}

type DeleteComment struct {
	CommentID int `json:"comment_id"`
}

func (d *DeleteComment) Parse(ctx *fiber.Ctx) *DeleteComment {
	var err error

	err = ctx.BodyParser(d)

	if err != nil {
		panic(cerrors.ErrInvalidRequest)
	}

	return d
}
