package rest

import (
	"github.com/gofiber/fiber/v2"
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendJsonResponse(ctx *fiber.Ctx, code int, res *JsonResponse) error {
	return ctx.Status(code).JSON(res)
}
