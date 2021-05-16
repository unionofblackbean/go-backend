package responses

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendEndpointNotFoundResponse(ctx *fiber.Ctx) {
	SendMessageResponse(ctx, http.StatusNotFound, "endpoint not found")
}
