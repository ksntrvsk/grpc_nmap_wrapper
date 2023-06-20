package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

// Types
type Config struct {
	Server struct {
		Host string `yml:"host"`
		Port string `yml:"port"`
	} `yml:"server"`
	Logger struct {
		Level string `yml:"level"`
	} `yml:"logger"`
}

// Constants
const path = "config.yml"

// Variables
var instance *Config
var once sync.Once

// Public methods
func NewCongif() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig(path, instance); err != nil {
			errDesc, _ := cleanenv.GetDescription(instance, nil)
			log.Fatal(errDesc)
		}
	})

	return instance, nil
}
