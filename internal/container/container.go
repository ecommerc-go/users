package container

import (
	"fmt"
	"log"

	"github.com/ecommerc-go/users/internal/config"
	"github.com/ecommerc-go/users/internal/lib/logger"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"github.com/ecommerc-go/users/internal/service"
	transport "github.com/ecommerc-go/users/internal/transport/grpc"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	Config *config.Config
	Api    *transport.Implementation
}

func NewContainer() *Container {

	// конфиг
	cfg := config.MustLoad()

	//подключение к БД
	connectCmd := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database, cfg.Postgres.Host, cfg.Postgres.Port)

	conn, err := sqlx.Connect("postgres", connectCmd)
	if err != nil {
		log.Fatal(err)
	}
	// инициализация логгера
	logger.InitLogger(cfg.Service.Mode)

	repos := repository.NewRepository(conn)
	serv := service.NewService(repos, cfg.JWTSecret.JWTSECRET)

	api := transport.NewImplementation(serv)

	return &Container{
		Config: cfg,
		Api:    api,
	}

}
