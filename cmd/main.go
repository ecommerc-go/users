package main

import (
	"fmt"
	"net"
	"os"

	"github.com/ecommerc-go/users/internal/container"
	"github.com/ecommerc-go/users/pkg/users"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {

	// Инициализация контейнера зависимостей
	container := container.NewContainer()

	// Создание нового gRPC-сервера
	grpcSrv := grpc.NewServer()

	// Регистрация сервиса пользователей в gRPC-сервере
	users.RegisterUserServiceServer(grpcSrv, container.Api)

	// Создание TCP-слушателя на порту, указанном в конфигурации
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", container.Config.Service.Port))
	if err != nil {
		slog.Error("failed to listen: %v", err)
		os.Exit(1)
	}

	// Запуск gRPC-сервера для обработки входящих запросов
	err = grpcSrv.Serve(lis)
	if err != nil {
		slog.Error("failed to serve: %v", err)
		os.Exit(1)
	}
}
