package config

import "sync"

import "github.com/BurntSushi/toml"

const (
	configPath = "./config/app.toml"
)

var (
	instance *Config
	once     sync.Once
)

// Config ..
type Config struct {
	Server ServerConfig
}

// NewConfig __constructor
func NewConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if _, err := toml.DecodeFile(configPath, instance); err != nil {
			panic(err)
		}
	})

	return instance
}
