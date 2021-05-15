package auth

type RestConfig struct {
	BindAddress string `json:"bind_addr"`
	BindPort    uint16 `json:"bind_port"`
}

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Address string `json:"addr"`
	Port    uint16 `json:"port"`

	Name string `json:"name"`
}

type Config struct {
	Rest     RestConfig     `json:"rest"`
	Database DatabaseConfig `json:"db"`
}
