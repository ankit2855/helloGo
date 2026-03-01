package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Port int
}

type Config struct {
	Env         string     `yaml:"env" env-required:"true" env-default:"dev"`
	StoragePath string     `yaml:"storagePath" env-required:"true"`
	HttpServer  HttpServer `yaml:"httpServer" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "config/local.yaml"
	}

	if configPath == "" {
		log.Fatalf("CONFIG_PATH is not set")
	}

	var config Config

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s does not exist", configPath)
	}

	err := cleanenv.ReadConfig(configPath, &config)

	if err != nil {
		log.Fatalf("Failed to read config: %v", err.Error())
	}

	return &config
}
