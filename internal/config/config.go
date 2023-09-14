package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"path"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env:"Env" default:"local"`
	StoragePath string     `yaml:"storage_path" default:"./storage/storage.db"`
	HttpServer  HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" default:"localhost:8080"`
	TimeOut     time.Duration `yaml:"timeout" default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" default:"60s""`
}

func MustLoad() *Config {
	pwd, _ := os.Getwd()
	configPath := path.Join(pwd, "config", "local.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("error in th config file, %s", err)
	}

	return &cfg
}
