package auth

import (
	"github.com/unionofblackbean/backend/pkg/database"
	"github.com/unionofblackbean/backend/pkg/rest"
)

type Config struct {
	Rest     rest.Config     `json:"rest"`
	Database database.Config `json:"db"`
}
