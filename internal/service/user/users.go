package service

import (
	"context"
	"errors"

	"github.com/ecommerc-go/users/internal/service"
)

type DbRepo interface {
	CreateUser(ctx context.Context, user *service.RegisterUserRequest) (*service.RegisterUserResponse, error)
	GetUser(ctx context.Context, id string) (*service.GetProfileResponse, error)
	DeleteUser(ctx context.Context, id string) error
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

func (r *Service) GetProfile(ctx context.Context, id string) (*service.GetProfileResponse, error) {
	data, _ := r.db.GetUser(ctx, id)

	return &service.GetProfileResponse{
		Profile: data.Profile,
	}, nil
}

func (r *Service) DeleteProfile(ctx context.Context, id string) error {
	err := r.db.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
