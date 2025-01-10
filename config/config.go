package config

import (
	"fmt"
)

// global config
type Config struct {
	Port string
	Name string
}

var config *Config

func InitConfig(name, port string) *Config {
	config = &Config{
		Name: name,
		Port: port,
	}
	fmt.Printf("%s config has been set!\n", name)
	return config
}

func GetConfig() *Config {
	return config
}
