package repository

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *RepoRegisterReq) (string, error)
	GetUser(ctx context.Context, id string) (*RepoUserProfile, error)
	DeleteUser(ctx context.Context, id string) error
}
