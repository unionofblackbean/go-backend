package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Pool struct {
	p *pgxpool.Pool
}

func (p *Pool) Exec(sql string, args ...interface{}) (err error) {
	_, err = p.p.Exec(context.Background(), sql, args)
	return
}

func (p *Pool) Query(sql string, args ...interface{}) (rows pgx.Rows, err error) {
	rows, err = p.p.Query(context.Background(), sql, args)
	return
}

func (p *Pool) QueryRow(sql string, args ...interface{}) pgx.Row {
	return p.p.QueryRow(context.Background(), sql, args)
}

func NewPool(
	username string, password string,
	addr string, port uint16,
	name string,
) (pool *Pool, err error) {
	pool = new(Pool)
	pool.p, err = pgxpool.Connect(context.Background(),
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s",
			username, password, addr, port, name,
		),
	)
	return
}
