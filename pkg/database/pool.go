package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Pool(
	username string, password string,
	addr string, port uint16,
	name string,
) (pool *pgxpool.Pool, err error) {
	pool, err = pgxpool.Connect(
		context.Background(),
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s",
			username, password, addr, port, name,
		),
	)
	return
}
