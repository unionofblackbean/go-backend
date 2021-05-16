package webhelpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/backend/pkg/responses"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	responses.SendInternalServerErrorReponse(ctx, err)
	return nil
}
