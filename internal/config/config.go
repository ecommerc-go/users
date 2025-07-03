package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres Postgres
	Service  Service
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Database string `env:"PG_DATABASE_NAME"`
}

type Service struct {
	Port string `env:"GRPC_PORT"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = ".env"
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %v", err)
	}
	fmt.Println(cfg)
	return &cfg

}
