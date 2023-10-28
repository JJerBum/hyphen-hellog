package model

import (
	"hyphen-hellog/cerrors"
	"hyphen-hellog/ent"
	"hyphen-hellog/verifier"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreateComment 클라이언트로부터 요청된 바디 값과 매핑되는 구조체 입니다.
// PostID는 Params() 함수를 이용해야 합니다해
// PostID와 ParentID는 BodyParser()함수를 이용해야 합니다.
type InCreateComment struct {
	PostID   int    `json:"post_id" validate:"required"`
	Content  string `json:"content" validate:"required"`
	ParentID int    `json:"parent_id" validate:"required"`
}

func (i *InCreateComment) ParseX(c *fiber.Ctx) *InCreateComment {
	var err error

	err = c.BodyParser(i)
	i.PostID, err = strconv.Atoi(c.Params("post_id"))

	if err != nil {
		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	verifier.Validate(i)

	return i
}

// GetComments 클라이언트로부터 요청된 바디 값과 매핑되는 구조체 입니다.
// PostID는 Params() 함수를 이용해야 합니다.
type InGetComments struct {
	PostID int
}

func (i *InGetComments) ParseX(c *fiber.Ctx) *InGetComments {
	var err error

	i.PostID, err = strconv.Atoi(c.Params("post_id"))

	if err != nil {
		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	return i
}

type InUpdateComment struct {
	CommentID int    `json:"comment_id" validate:"required"`
	Content   string `json:"content" validate:"required"`
}

func (i *InUpdateComment) ParseX(c *fiber.Ctx) *InUpdateComment {
	var err error

	err = c.BodyParser(i)

	if err != nil {
		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	verifier.Validate(i)

	return i
}

type InDeleteComment struct {
	CommentID int `json:"comment_id" validate:"required"`
}

func (i *InDeleteComment) ParseX(c *fiber.Ctx) *InDeleteComment {
	var err error

	err = c.BodyParser(i)

	if err != nil {
		panic(cerrors.ParsingErr{
			Err: err.Error(),
		})
	}

	verifier.Validate(i)

	return i
}

// response
type OutGetComments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	*ent.Comment     `json:"comment"`
	*ent.Author      `json:"author"`
	CommentOfComment []CommentOfComment `json:"comment_of_comment"`
}

type CommentOfComment struct {
	*ent.Comment `json:"comment"`
	*ent.Author  `json:"author"`
}
