package rest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SendDataResponse(ctx *fiber.Ctx, data interface{}) (err error) {
	err = SendJsonResponse(ctx, http.StatusOK, &JsonResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
	if err != nil {
		err = fmt.Errorf("failed to send JSON response -> %v", err)
	}
	return
}
