package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendResourceAlreadyExistsResponse(ctx *fiber.Ctx) {
	SendMessageResponse(ctx, http.StatusConflict, "resource already exists")
}
