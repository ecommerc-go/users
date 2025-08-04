package service

import (
	"context"

	"github.com/ecommerc-go/users/internal/models"
)

//go:generate minimock -i github.com/ecommerc-go/users/internal/service.UserService -o ./mocks -s "_minimock.go"

type UserService interface {
	RegisterUser(ctx context.Context, req *models.RegisterRequest) (string, error)
	LoginUser(ctx context.Context, req *models.LoginUserRequest) (string, error)
	GetProfile(ctx context.Context, id string) (*models.UserProfile, error)
	UpdateProfile(ctx context.Context, req *models.UpdateProfileRequest) error
	DeleteProfile(ctx context.Context, id string) error
}
