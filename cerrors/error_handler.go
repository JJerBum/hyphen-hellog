package cerrors

import (
	"hyphen-hellog/model/response"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	return ctx.Status(fiber.StatusInternalServerError).JSON(response.Genreal{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}
