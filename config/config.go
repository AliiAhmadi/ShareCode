package config

import (
	"strconv"
)

type Config struct {
	Port    int
	Address string
}

func (config *Config) Get() string {
	return config.Address + ":" + strconv.Itoa(config.Port)
}
