package database

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Address string `json:"addr"`
	Port    uint16 `json:"port"`

	Name string `json:"name"`
}
