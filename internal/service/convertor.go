package service

import (
	"github.com/ecommerc-go/users/internal/repository"
)

func RegistrationToRepo(data *SvcRegisterRequest) *repository.RepoRegisterReq {
	return &repository.RepoRegisterReq{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
	}
}

func UserProfileToSVC(data *repository.RepoUserProfile) *SvcUserProfile {
	return &SvcUserProfile{
		Name:    data.Name,
		Email:   data.Email,
		Address: data.Address,
	}
}
