package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendUnsupportedMethodResponse(ctx *fiber.Ctx) {
	SendMessageResponse(ctx, http.StatusMethodNotAllowed, "unsupported method")
}
