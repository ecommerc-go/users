package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ecommerc-go/users/internal/container"
	"github.com/ecommerc-go/users/pkg/users"
	"google.golang.org/grpc"
)

func main() {

	container := container.NewContainer()

	grpcSrv := grpc.NewServer()
	users.RegisterUserServiceServer(grpcSrv, container.Api)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", container.Config.Service.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = grpcSrv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
