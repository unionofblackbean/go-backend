package rest

type Config struct {
	BindAddress string `json:"bind_addr"`
	BindPort    uint16 `json:"bind_port"`
}
