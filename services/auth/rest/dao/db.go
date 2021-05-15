package dao

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func Init(p *pgxpool.Pool) {
	pool = p
}

func checkPool() (err error) {
	if pool == nil {
		err = errors.New("dao database connection pool not initialized")
	}
	return
}
