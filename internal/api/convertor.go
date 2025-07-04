package api

import (
	"github.com/ecommerc-go/users/internal/service"
	"github.com/ecommerc-go/users/pkg/users"
)

func RegisterUserFromProto(data *users.RegisterUserRequest) *service.SvcRegisterRequest {
	return &service.SvcRegisterRequest{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
	}
}

func RegisterUserToProto(data *service.SvcRegisterRequest) *users.RegisterUserRequest {
	return &users.RegisterUserRequest{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
	}
}

func LoginUserFromProto(data *users.LoginUserRequest) *service.LoginUserRequest {
	return &service.LoginUserRequest{
		Email:    data.Email,
		Password: data.Password,
	}
}

func UserProfileToProto(data *service.SvcUserProfile) *users.GetProfileResponse {
	return &users.GetProfileResponse{
		Profile: &users.UserProfile{
			UserId:    data.Id,
			Email:     data.Email,
			Name:      data.Name,
			Address:   data.Address,
			CreatedAt: data.Created_at,
		},
	}
}

func UpdateProfileFromProto(data *users.UpdateProfileRequest) *service.UpdateProfileRequest {
	return &service.UpdateProfileRequest{
		Id:      data.UserId,
		Name:    data.Name,
		Address: data.Address,
	}
}
