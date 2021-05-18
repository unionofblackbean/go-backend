package rest

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	SendInternalServerErrorReponse(ctx, err)
	return nil
}
