package model

import (
	"hyphen-hellog/cerrors"
	"hyphen-hellog/ent"
	"hyphen-hellog/verifier"
	"mime/multipart"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreatePost는 클라이언트로부터 요청된 바디 값과 매핑되는 구조체 입니다.
// Title과 Content는 BodyParser()함수를 이용해야 합니다.
// PreviwImage는 FormFile() 함수를 이용해야 합니다.
// IsPrivate는 FormValue() 함수를 이용해야 합니다.
type InCreatePost struct {
	Title            string                `form:"title"  validate:"required"`
	Content          string                `form:"content" validate:"required"`
	PreviewImage     *multipart.FileHeader `form:"preview_image" validate:"required"`
	ShortDescription string                `form:"short_description"`
	IsPrivate        bool                  `form:"is_private"`
}

func (i *InCreatePost) ParseX(c *fiber.Ctx) *InCreatePost {
	var err error

	err = c.BodyParser(i)
	i.IsPrivate = c.FormValue("is_private") == "true"
	i.PreviewImage, err = c.FormFile("image")
	if err != nil {

		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	verifier.Validate(i)

	return i
}

type InGetPost struct {
	PostID int `json:"post_id" validate:"required"`
}

func (i *InGetPost) ParseX(c *fiber.Ctx) *InGetPost {
	var err error

	i.PostID, err = strconv.Atoi(c.Params("post_id"))
	if err != nil {
		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	verifier.Validate(i)

	return i
}

type InUpdatePost struct {
	Title        string                `form:"title"  validate:"required"`
	Content      string                `form:"content" validate:"required"`
	PreviewImage *multipart.FileHeader `form:"preview_image" validate:"required"`
	IsPrivate    bool                  `form:"is_private" validate:"boolean"`
	PostID       int                   `json:"post_id"`
}

func (i *InUpdatePost) ParseX(c *fiber.Ctx) *InUpdatePost {
	var err error

	err = c.BodyParser(i)
	i.PostID, err = strconv.Atoi(c.Params("post_id"))
	i.PreviewImage, err = c.FormFile("preview_image")

	if err != nil {
		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	verifier.Validate(i)

	return i
}

type InDeletePost struct {
	PostID int `validate:"required"`
}

func (i *InDeletePost) ParseX(c *fiber.Ctx) *InDeletePost {
	var err error

	i.PostID, err = strconv.Atoi(c.Params("post_id"))

	if err != nil {
		panic(cerrors.DeleteErr{
			Err: err.Error(),
		})
	}

	return i
}

type OutGetPost struct {
	*ent.Post   `json:"post"`
	IsLiked     bool `json:"is_liked"`
	*ent.Author `json:"author"`
}

type OutGetPosts struct {
	Posts []*OutGetPost `json:"posts"`
}
