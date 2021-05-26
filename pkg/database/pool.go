package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Pool struct {
	p *pgxpool.Pool
}

func NewPool(
	username string, password string,
	addr string, port uint16,
	name string,
) (*Pool, error) {
	var err error
	pool := new(Pool)
	pool.p, err = pgxpool.Connect(context.Background(),
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s",
			username, password, addr, port, name,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection with database -> %v", err)
	}

	return pool, nil
}

func (p *Pool) Validate() (err error) {
	if p == nil {
		err = errors.New("database connection pool not initialized")
	}

	return
}

func (p *Pool) Exec(sql string, args ...interface{}) (err error) {
	_, err = p.p.Exec(context.Background(), sql, args)
	if err != nil {
		err = fmt.Errorf("failed to execute SQL -> %v", err)
	}

	return
}

func (p *Pool) Query(sql string, args ...interface{}) (pgx.Rows, error) {
	rows, err := p.p.Query(context.Background(), sql, args)
	if err != nil {
		return nil, fmt.Errorf("failed to query rows from database -> %v", err)
	}

	return rows, nil
}

func (p *Pool) QueryRow(sql string, args ...interface{}) pgx.Row {
	return p.p.QueryRow(context.Background(), sql, args)
}
