package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/backend/pkg/database"
	"github.com/unionofblackbean/backend/pkg/rest"
	"github.com/unionofblackbean/backend/services/auth/api/controllers"
	"github.com/unionofblackbean/backend/services/auth/api/dao"
)

func Init(pool *database.Pool) {
	dao.Init(pool)
}

func Run(addr string, port uint16) error {
	r := fiber.New(fiber.Config{
		ErrorHandler: rest.ErrorHandler,
	})

	r.All("/users/:uuid?", controllers.Users)

	r.Use(rest.NotFoundMiddleWare)

	return r.Listen(fmt.Sprintf("%s:%d", addr, port))
}
