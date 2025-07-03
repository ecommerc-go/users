package container

import (
	"fmt"
	"log"

	"github.com/ecommerc-go/users/internal/api"
	"github.com/ecommerc-go/users/internal/config"
	"github.com/ecommerc-go/users/internal/repository"
	service "github.com/ecommerc-go/users/internal/service/user"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	Config *config.Config
	Api    *api.Implementation
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

	repos := repository.NewRepository(conn)
	serv := service.NewService(repos)

	api := api.NewImplementation(serv)

	return &Container{
		Config: cfg,
		Api:    api,
	}

}
