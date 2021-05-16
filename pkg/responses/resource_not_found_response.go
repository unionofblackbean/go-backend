package responses

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendResourceNotFoundResponse(ctx *fiber.Ctx) {
	SendMessageResponse(ctx, http.StatusNotFound, "resource not found")
}
