package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool = nil

func Pool() (*pgxpool.Pool, error) {
	if pool == nil {
		return nil, errors.New("database connection pool not initialized")
	}

	return pool, nil
}

func InitPool(addr string, port uint16, username string, password string, name string) (err error) {
	pool, err = pgxpool.Connect(
		context.Background(),
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s",
			username, password, addr, port, name,
		),
	)
	return
}
