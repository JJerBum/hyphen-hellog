package main

import (
	"encoding/json"
	"flag"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/cerrors/exception"
	"hyphen-hellog/controller"
	"hyphen-hellog/initializer"
	"hyphen-hellog/model"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var port string

func init() {
	flagPort := flag.Int("p", 8080, "Enter the port")
	flag.Parse()
	port = ":" + strconv.Itoa(*flagPort)

	// load env variables
	initializer.LoadEnv()
}

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})
	app.Use(recover.New())

	log.Fatal(controller.Route(app).Listen(port))

}

func errorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(cerrors.ValidationErr)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		exception.Sniff(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.General{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	if _, createErr := err.(cerrors.CreateErr); createErr {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.General{
			Code:    500,
			Message: "Not Created",
			Data:    err.Error(),
		})
	}

	if _, selectErr := err.(cerrors.SelectErr); selectErr {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.General{
			Code:    500,
			Message: "Not Selected",
			Data:    err.Error(),
		})
	}

	if _, updateErr := err.(cerrors.UpdateErr); updateErr {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.General{
			Code:    500,
			Message: "Not Updated",
			Data:    err.Error(),
		})
	}

	if _, deleteErr := err.(cerrors.UpdateErr); deleteErr {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.General{
			Code:    500,
			Message: "Not Deleted",
			Data:    err.Error(),
		})
	}

	if _, parsingErr := err.(cerrors.ParsingErr); parsingErr {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.General{
			Code:    400,
			Message: "Don't Parsing",
			Data:    err.Error(),
		})
	}

	if _, requestFailedErr := err.(cerrors.RequestFailedErr); requestFailedErr {
		return ctx.Status(fiber.StatusBadGateway).JSON(model.General{
			Code:    502,
			Message: "Reqeust Failed",
			Data:    err.Error(),
		})
	}

	if _, wrongApproachErr := err.(cerrors.WrongApproachErr); wrongApproachErr {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.General{
			Code:    401,
			Message: "Wrong Approach",
			Data:    err.Error(),
		})
	}

	if _, unauthorizedErr := err.(cerrors.UnauthorizedErr); unauthorizedErr {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.General{
			Code:    401,
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.General{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}
