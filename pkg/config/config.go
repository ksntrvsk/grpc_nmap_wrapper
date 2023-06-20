package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Server struct {
		Host string `yml:"host"`
		Port string `yml:"port"`
	} `yml:"server"`
	Logger struct {
		Level string `yml:"level"`
	} `yml:"logger"`
}

var instance *Config
var once sync.Once

func NewCongif() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			errDesc, _ := cleanenv.GetDescription(instance, nil)
			log.Fatal(errDesc)
		}
	})

	// Create config file without library

	// file, err := os.Open("config.yml")
	// if err != nil {
	// 	return nil, fmt.Errorf("unable to open config file: %v", err)
	// }
	// defer file.Close()

	// decoder := yaml.NewDecoder(file)
	// if err := decoder.Decode(instance); err != nil {
	// 	return nil, fmt.Errorf("unable to decode config file: %v", err)
	// }

	return instance, nil
}
