package service

import (
	"context"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *SvcRegisterRequest) (string, error)
	// LoginUser(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error)
	GetProfile(ctx context.Context, id string) (*SvcUserProfile, error)
	// UpdateProfile(ctx context.Context, req *UpdateProfileRequest) error
	DeleteProfile(ctx context.Context, id string) error
}
