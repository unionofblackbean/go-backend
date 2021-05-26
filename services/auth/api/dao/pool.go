package dao

import (
	"github.com/unionofblackbean/backend/pkg/database"
)

var pool *database.Pool

func Init(p *database.Pool) {
	pool = p
}
