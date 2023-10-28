package middleware

import (
	"hyphen-hellog/cerrors"
	"hyphen-hellog/client/user"
	"hyphen-hellog/database"
	"hyphen-hellog/ent"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error {

	// 검증된 유저 인가?
	response, err := user.Validate(c.Get("Authorization"))
	if err != nil {
		panic(cerrors.RequestFailedErr{
			Err: err.Error(),
		})
	}

	// 이미 있는 사용자 인가?
	author, err := database.Get().GetAuthorByAuthorID(c.Context(), response.Data)

	// 없는 유저라면
	if err != nil {
		// 사용자 등록하기
		author = database.Get().CreateAuthorX(c.Context(), &ent.Author{AuthorID: response.Data})
	}

	// local로 저장
	c.Locals("user", author)

	return c.Next()
}

func Auth(c *fiber.Ctx) error {
	var author *ent.Author = nil

	// 검증된 유저 인가?
	response, err := user.Validate(c.Get("Authorization"))
	if err == nil {
		// 이미 있는 사용자 인가?
		author, _ = database.Get().GetAuthorByAuthorID(c.Context(), response.Data)
	}

	c.Locals("user", author)

	return c.Next()
}
