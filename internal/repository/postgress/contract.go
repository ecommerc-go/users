package repository

import (
	"context"

	"github.com/ecommerc-go/users/internal/domain"
)

//go:generate minimock -i github.com/ecommerc-go/users/internal/repository/postgress.UserRepository -o ./mocks -s "_minimock.go"

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.RegisterUser) (string, error)
	GetUser(ctx context.Context, id string) (*domain.UserProfile, error)
	DeleteUser(ctx context.Context, id string) error
	GetCredentials(ctx context.Context, email string) (*domain.Creds, error)
	UpdateUser(ctx context.Context, user *domain.UpdateProfile) (int64, error)
}
