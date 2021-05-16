package smartbots

import "github.com/unionofblackbean/backend/pkg/database"

type RestConfig struct {
	BindAddress string `json:"bind_addr"`
	BindPort    uint16 `json:"bind_port"`
}

type Config struct {
	Rest     RestConfig      `json:"rest"`
	Database database.Config `json:"db"`
}
