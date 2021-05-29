package rest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendInternalServerErrorResponse(ctx *fiber.Ctx, err error) error {
	message := "an error occurred while processing the request"
	if err != nil {
		message = err.Error()
	}

	sendMessageErr := SendMessageResponse(ctx, http.StatusInternalServerError, message)
	if sendMessageErr != nil {
		return fmt.Errorf("failed to send message response -> %v", sendMessageErr)
	}
	return nil
}
