package service

import (
	"context"
	"errors"

	"github.com/ecommerc-go/users/internal/service"
)

type DbRepo interface {
	CreateUser(ctx context.Context, user *service.RegisterUserRequest) (*service.RegisterUserResponse, error)
}

var ErrUserNotFound = errors.New("user not found")

type Service struct {
	db DbRepo
}

func NewService(db DbRepo) *Service {
	return &Service{
		db: db,
	}
}

func (r *Service) RegisterUser(ctx context.Context, req *service.RegisterUserRequest) (*service.RegisterUserResponse, error) {
	data, _ := r.db.CreateUser(ctx, req)

	return &service.RegisterUserResponse{
		Id: data.Id,
	}, nil
}
