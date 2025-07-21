package repository

import (
	"context"

	"github.com/ecommerc-go/users/internal/models"
)

//go:generate minimock -i github.com/ecommerc-go/users/internal/repository/postgress.UserRepository -o ./mocks -s "_minimock.go"

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.RegisterRequest) (string, error)
	GetUser(ctx context.Context, id string) (*models.UserProfile, error)
	DeleteUser(ctx context.Context, id string) error
	GetCredentials(ctx context.Context, email string) (*Creds, error)
	UpdateUser(ctx context.Context, user *models.UpdateProfileRequest) (int64, error)
}
