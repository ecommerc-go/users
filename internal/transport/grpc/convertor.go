package transport

import (
	"github.com/ecommerc-go/users/internal/domain"
	"github.com/ecommerc-go/users/pkg/users"
)

func RegisterUserFromProto(data *users.RegisterUserRequest) *domain.RegisterUser {
	return &domain.RegisterUser{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
	}
}

func RegisterUserToProto(data *domain.RegisterUser) *users.RegisterUserRequest {
	return &users.RegisterUserRequest{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
	}
}

func LoginUserFromProto(data *users.LoginUserRequest) *domain.LoginUser {
	return &domain.LoginUser{
		Email:    data.Email,
		Password: data.Password,
	}
}

func UserProfileToProto(data *domain.UserProfile) *users.GetProfileResponse {
	return &users.GetProfileResponse{
		Profile: &users.UserProfile{
			UserId:    data.ID,
			Email:     data.Email,
			Name:      data.Name,
			Address:   data.Address,
			CreatedAt: data.Created_at,
		},
	}
}

func UpdateProfileFromProto(data *users.UpdateProfileRequest) *domain.UpdateProfile {
	return &domain.UpdateProfile{
		ID:      data.UserId,
		Name:    data.Name,
		Address: data.Address,
	}
}
