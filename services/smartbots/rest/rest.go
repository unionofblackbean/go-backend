package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/backend/pkg/database"
	"github.com/unionofblackbean/backend/pkg/webhelpers"
	"github.com/unionofblackbean/backend/services/auth/rest/dao"
)

func Init(pool *database.Pool) {
	dao.Init(pool)
}

func Run(addr string, port uint16) error {
	r := fiber.New(fiber.Config{
		ErrorHandler: webhelpers.ErrorHandler,
	})

	r.Use(webhelpers.NotFoundMiddleWare)

	return r.Listen(fmt.Sprintf("%s:%d", addr, port))
}
