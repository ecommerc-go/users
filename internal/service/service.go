package service

import (
	"context"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error)
	// LoginUser(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error)
	GetProfile(ctx context.Context, id string) (*GetProfileResponse, error)
	// UpdateProfile(ctx context.Context, req *UpdateProfileRequest) error
	DeleteProfile(ctx context.Context, id string) error
}
