package rest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendUnsupportedMethodResponse(ctx *fiber.Ctx) (err error) {
	err = SendMessageResponse(ctx, http.StatusMethodNotAllowed, "unsupported method")
	if err != nil {
		err = fmt.Errorf("failed to send message response -> %v", err)
	}
	return
}
