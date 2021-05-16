package dao

import (
	"errors"

	"github.com/unionofblackbean/backend/pkg/database"
)

var pool *database.Pool

func Init(p *database.Pool) {
	pool = p
}

func checkPool() (err error) {
	if pool == nil {
		err = errors.New("dao database connection pool not initialized")
	}
	return
}
