package main

import (
	"log"

	"github.com/unionofblackbean/backend/pkg/database"
	"github.com/unionofblackbean/backend/services/auth/rest"
)

func main() {
	pool, err := database.Pool(
		"backend", "backend",
		"127.0.0.1", 5432,
		"backend",
	)
	if err != nil {
		log.Fatalf("failed to establish connection with database -> %v", err)
	}

	rest.Init(pool)
	rest.Run("127.0.0.1", 8080)
}
