package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/backend/pkg/responses"
)

func NotFound(ctx *fiber.Ctx) error {
	responses.SendNotFoundResponse(ctx)
	return nil
}
