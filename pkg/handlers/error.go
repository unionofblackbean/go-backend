package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/backend/pkg/responses"
)

func Error(ctx *fiber.Ctx, err error) error {
	responses.SendInternalServerErrorReponse(ctx, err)
	return nil
}
