package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config структура для конфига
type Config struct {
	Postgres  Postgres
	Service   Service
	JWTSecret JWTSECRET
}

// Postgres содержит поля конфигурации для подключения к базе данных PostgreSQL.
type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Database string `env:"PG_DATABASE_NAME"`
}

type JWTSECRET struct {
	JWTSECRET string `env:"JWT_SECRET"`
}

type Service struct {
	Port string `env:"GRPC_PORT"`
	Mode string `env:"MODE"`
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
