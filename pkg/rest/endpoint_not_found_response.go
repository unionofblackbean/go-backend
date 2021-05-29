package rest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendEndpointNotFoundResponse(ctx *fiber.Ctx) (err error) {
	err = SendMessageResponse(ctx, http.StatusNotFound, "endpoint not found")
	if err != nil {
		err = fmt.Errorf("failed to send message response -> %v", err)
	}
	return
}
