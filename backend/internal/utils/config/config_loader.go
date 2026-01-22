package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port          string `yaml:"port"`
		AllowedOrigin string `yaml:"allowed_origin"`
	} `yaml:"server"`
	Database struct {
		Driver string `yaml:"driver"`
		DSN    string `yaml:"dsn"`
	} `yaml:"database"`
	Session struct {
		SecretKey string `yaml:"secret_key"`
		MaxAge    int    `yaml:"max_age"`
	} `yaml:"session"`
}

var AppConfig Config

func LoadConfig() error {
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, &AppConfig)
}