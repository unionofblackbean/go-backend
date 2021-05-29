package rest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendResourceAlreadyExistsResponse(ctx *fiber.Ctx) (err error) {
	err = SendMessageResponse(ctx, http.StatusConflict, "resource already exists")
	if err != nil {
		err = fmt.Errorf("failed to send message response -> %v", err)
	}
	return
}
