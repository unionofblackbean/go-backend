package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/backend/pkg/responses"
)

func NotFound(ctx *fiber.Ctx) error {
	responses.SendEndpointNotFoundResponse(ctx)
	return nil
}
