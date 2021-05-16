package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unionofblackbean/backend/pkg/database"
	"github.com/unionofblackbean/backend/pkg/handlers"
	"github.com/unionofblackbean/backend/services/auth/rest/controllers"
	"github.com/unionofblackbean/backend/services/auth/rest/dao"
)

func Init(pool *database.Pool) {
	dao.Init(pool)
}

func Run(addr string, port uint16) error {
	r := fiber.New(fiber.Config{
		ErrorHandler: handlers.Error,
	})

	r.All("/users/:uuid?", controllers.Users)

	r.Use(handlers.NotFound)

	return r.Listen(fmt.Sprintf("%s:%d", addr, port))
}
