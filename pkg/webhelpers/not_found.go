package webhelpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/backend/pkg/responses"
)

func NotFoundMiddleWare(ctx *fiber.Ctx) error {
	responses.SendEndpointNotFoundResponse(ctx)
	return nil
}
