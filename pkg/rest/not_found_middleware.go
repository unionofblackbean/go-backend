package rest

import (
	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleWare(ctx *fiber.Ctx) error {
	return SendEndpointNotFoundResponse(ctx)
}
