package apiserver

import "github.com/Pupye/rest-api/internal/app/store"

//Config ...
type Config struct {
	BindAddr string        `toml:"bind_addr"`
	LogLevel string        `toml:"log_level"`
	Store    *store.Config `toml:"database_configs"`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
