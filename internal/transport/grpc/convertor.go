package transport

import (
	"github.com/ecommerc-go/users/internal/models"
	"github.com/ecommerc-go/users/pkg/users"
)

func RegisterUserFromProto(data *users.RegisterUserRequest) *models.RegisterRequest {
	return &models.RegisterRequest{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
	}
}

func RegisterUserToProto(data *models.RegisterRequest) *users.RegisterUserRequest {
	return &users.RegisterUserRequest{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
	}
}

func LoginUserFromProto(data *users.LoginUserRequest) *models.LoginUserRequest {
	return &models.LoginUserRequest{
		Email:    data.Email,
		Password: data.Password,
	}
}

func UserProfileToProto(data *models.UserProfile) *users.GetProfileResponse {
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

func UpdateProfileFromProto(data *users.UpdateProfileRequest) *models.UpdateProfileRequest {
	return &models.UpdateProfileRequest{
		ID:      data.UserId,
		Name:    data.Name,
		Address: data.Address,
	}
}
