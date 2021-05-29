package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendInternalServerErrorResponse(ctx *fiber.Ctx, err error) {
	message := "an error occurred while processing the request"
	if err != nil {
		message = err.Error()
	}
	SendMessageResponse(ctx, http.StatusInternalServerError, message)
}
