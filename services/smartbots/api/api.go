package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/go-backend/pkg/database"
	"github.com/unionofblackbean/go-backend/pkg/rest"
	"github.com/unionofblackbean/go-backend/services/smartbots/api/dao"
)

func Init(pool *database.Pool) {
	dao.Init(pool)
}

func Run(addr string, port uint16) error {
	r := fiber.New(fiber.Config{
		ErrorHandler: rest.ErrorHandler,
	})

	r.Use(rest.NotFoundMiddleWare)

	return r.Listen(fmt.Sprintf("%s:%d", addr, port))
}
