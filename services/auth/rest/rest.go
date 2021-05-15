package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/unionofblackbean/backend/pkg/responses"
	"github.com/unionofblackbean/backend/services/auth/rest/controllers"
	"github.com/unionofblackbean/backend/services/auth/rest/dao"
)

func Init(pool *pgxpool.Pool) {
	dao.Init(pool)
}

func Run(addr string, port uint16) error {
	r := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			responses.SendInternalServerErrorReponse(ctx, err)
			return nil
		},
	})

	r.All("/users/:uuid?", controllers.Users)

	r.Use(controllers.NotFound)

	return r.Listen(fmt.Sprintf("%s:%d", addr, port))
}
