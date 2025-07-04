package service

import (
	"context"
	"errors"

	"github.com/ecommerc-go/users/internal/repository"
)

var ErrUserNotFound = errors.New("user not found")

type Service struct {
	db repository.UserRepository
}

func NewService(db repository.UserRepository) *Service {
	return &Service{
		db: db,
	}
}

func (r *Service) RegisterUser(ctx context.Context, req *SvcRegisterRequest) (string, error) {
	newReq := RegistrationToRepo(req)
	data, err := r.db.CreateUser(ctx, newReq)
	if err != nil {
		///обрабтка и логи
	}

	return data, nil
}

func (r *Service) GetProfile(ctx context.Context, id string) (*SvcUserProfile, error) {
	data, _ := r.db.GetUser(ctx, id)
	convertData := UserProfileToSVC(data)

	return convertData, nil
}

func (r *Service) DeleteProfile(ctx context.Context, id string) error {
	err := r.db.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
