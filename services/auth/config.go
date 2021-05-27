package auth

import (
	"github.com/unionofblackbean/go-backend/pkg/database"
	"github.com/unionofblackbean/go-backend/pkg/rest"
)

type Config struct {
	Rest     rest.Config     `json:"rest"`
	Database database.Config `json:"db"`
}
