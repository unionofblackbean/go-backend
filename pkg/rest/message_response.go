package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func SendMessageResponse(ctx *fiber.Ctx, code int, message string) (err error) {
	err = SendJsonResponse(ctx, code, &JsonResponse{
		Code:    code,
		Message: message,
	})
	if err != nil {
		err = fmt.Errorf("failed to send JSON response -> %v", err)
	}
	return
}
