package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendDataResponse(ctx *fiber.Ctx, data interface{}) {
	ctx.Status(http.StatusOK).JSON(&JsonResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
