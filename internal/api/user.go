package api

import (
	"context"
	"fmt"

	"github.com/ecommerc-go/users/internal/service"
	"github.com/ecommerc-go/users/pkg/users"
)

type Implementation struct {
	users.UnimplementedUserServiceServer
	UserService service.UserService
}

func NewImplementation(UserService service.UserService) *Implementation {
	return &Implementation{
		UserService: UserService,
	}
}

func (s *Implementation) RegisterUser(ctx context.Context, req *users.RegisterUserRequest) (*users.RegisterUserResponse, error) {
	id, err := s.UserService.RegisterUser(ctx, RegisterUserFromProto(req))
	if err != nil {
		fmt.Println("Ошибка")
	}
	return &users.RegisterUserResponse{
		UserId: id,
	}, nil
}

// func (s *Implementation) LoginUser(ctx context.Context, req *users.LoginUserRequest) (*users.LoginUserResponse, error) {
// 	data, err := s.UserService.LoginUser(ctx, LoginUserFromProto(req))
// 	if err != nil {
// 		fmt.Println("Ошибка")
// 	}

// 	return &users.LoginUserResponse{
// 		AccessToken: data.Access_token,
// 	}, nil
// }

func (s *Implementation) GetProfile(ctx context.Context, req *users.GetProfileRequest) (*users.GetProfileResponse, error) {
	data, err := s.UserService.GetProfile(ctx, req.UserId)
	if err != nil {
		fmt.Println("Ошибка")
	}

	return UserProfileToProto(data), nil
}

// func (s *Implementation) UpdateProfile(ctx context.Context, req *users.UpdateProfileRequest) (*users.UpdateProfileResponse, error) {
// 	err := s.UserService.UpdateProfile(ctx, UpdateProfileFromProto(req))
// 	if err != nil {
// 		return &users.UpdateProfileResponse{
// 			Success: false,
// 		}, err
// 	}

// 	return &users.UpdateProfileResponse{
// 		Success: true,
// 	}, nil
// }

func (s *Implementation) DeleteProfile(ctx context.Context, req *users.DeleteProfileRequest) (*users.DeleteProfileResponse, error) {
	err := s.UserService.DeleteProfile(ctx, req.UserId)
	if err != nil {
		return &users.DeleteProfileResponse{
			Success: false,
		}, err
	}

	return &users.DeleteProfileResponse{
		Success: true,
	}, nil
}
