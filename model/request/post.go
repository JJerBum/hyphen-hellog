package request

import (
	"hyphen-hellog/cerrors"
	"mime/multipart"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreatePost는 클라이언트로부터 요청된 바디 값과 매핑되는 구조체 입니다.
// Title과 Content는 BodyParser()함수를 이용해야 합니다.
// PreviwImage는 FormFile() 함수를 이용해야 합니다.
// IsPrivate는 FormValue() 함수를 이용해야 합니다.
type CreatePost struct {
	Title        string                `form:"title"  validate:"required"`
	Content      string                `form:"content" validate:"required"`
	PreviewImage *multipart.FileHeader `form:"preview_image" validate:"required"`
	IsPrivate    bool                  `form:"is_private" validate:"boolean"`
}

func (c *CreatePost) Parse(ctx *fiber.Ctx) *CreatePost {
	var err error

	err = ctx.BodyParser(c)
	c.IsPrivate = ctx.FormValue("is_private") == "true"
	c.PreviewImage, err = ctx.FormFile("preview_image")

	if err != nil {
		panic(cerrors.ErrInvalidRequest)
	}

	return c
}

type GetPost struct {
	PostID int `json:"post_id"`
}

func (c *GetPost) Parse(ctx *fiber.Ctx) *GetPost {
	var err error

	c.PostID, err = strconv.Atoi(ctx.Params("post_id"))
	if err != nil {
		panic(cerrors.ErrInvalidRequest)
	}

	return c
}

type UpdatePost struct {
	Title        string                `form:"title"  validate:"required"`
	Content      string                `form:"content" validate:"required"`
	PreviewImage *multipart.FileHeader `form:"preview_image" validate:"required"`
	IsPrivate    bool                  `form:"is_private" validate:"boolean"`
	PostID       int                   `json:"post_id"`
}

func (u *UpdatePost) Parse(ctx *fiber.Ctx) *UpdatePost {
	var err error

	err = ctx.BodyParser(u)
	u.PostID, err = strconv.Atoi(ctx.Params("post_id"))
	u.PreviewImage, err = ctx.FormFile("preview_image")

	if err != nil {
		panic(cerrors.ErrInvalidRequest)
	}

	return u
}
