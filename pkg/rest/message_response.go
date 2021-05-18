package rest

import "github.com/gofiber/fiber/v2"

func SendMessageResponse(ctx *fiber.Ctx, code int, message string) {
	ctx.Status(code).JSON(&JsonResponse{
		Code:    code,
		Message: message,
	})
}
