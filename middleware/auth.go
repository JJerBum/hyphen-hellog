package middleware

import (
	"hyphen-hellog/client/user"
	"hyphen-hellog/database"
	"hyphen-hellog/ent"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error {

	// 검증된 유저 인가?
	response := user.ValidateX(c.Get("Authorization"))

	// 이미 있는 사용자 인가?
	author, err := database.New().GetAuthorByAuthorID(c.Context(), response.Data)

	// 없는 유저라면
	if err != nil {
		// 사용자 등록하기
		author = database.New().CreateAuthorX(c.Context(), &ent.Author{AuthorID: response.Data})
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
		author, _ = database.New().GetAuthorByAuthorID(c.Context(), response.Data)
	}

	c.Locals("user", author)

	return c.Next()
}
