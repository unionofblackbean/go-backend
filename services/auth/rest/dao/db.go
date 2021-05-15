package dao

import "github.com/jackc/pgx/v4/pgxpool"

var pool *pgxpool.Pool

func Init(p *pgxpool.Pool) {
	pool = p
}
