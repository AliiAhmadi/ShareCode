package config

import (
	"strconv"
)

type Config struct {
	Port    int
	Address string
	DSN     string
	Secret  string
}

func (config *Config) Get() string {
	return config.Address + ":" + strconv.Itoa(config.Port)
}
