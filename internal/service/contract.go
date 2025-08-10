package service

import (
	"context"

	"github.com/ecommerc-go/users/internal/domain"
)

//go:generate minimock -i github.com/ecommerc-go/users/internal/service.UserService -o ./mocks -s "_minimock.go"

type UserService interface {
	RegisterUser(ctx context.Context, req *domain.RegisterUser) (string, error)
	LoginUser(ctx context.Context, req *domain.LoginUser) (string, error)
	GetProfile(ctx context.Context, id string) (*domain.UserProfile, error)
	UpdateProfile(ctx context.Context, req *domain.UpdateProfile) error
	DeleteProfile(ctx context.Context, id string) error
}
