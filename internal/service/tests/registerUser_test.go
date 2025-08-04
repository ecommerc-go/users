package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/ecommerc-go/users/internal/models"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	repoMocks "github.com/ecommerc-go/users/internal/repository/postgress/mocks"
	"github.com/ecommerc-go/users/internal/service"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterUser(t *testing.T) {

	type UserRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *models.RegisterRequest
	}

	var (
		ctx    = context.Background()
		mc     = minimock.NewController(t)
		userID = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
		req    = &models.RegisterRequest{
			Email:    "den@yandex.ru",
			Password: "12345678",
			Name:     "Denis",
			Address:  "Марс",
		}
	)

	tests := []struct {
		name               string
		args               args
		want               string
		err                error
		userRepositoryMock UserRepositoryMockFunc
	}{
		{
			name: "register success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: userID,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Set(func(ctx context.Context, req *models.RegisterRequest) (string, error) {
					require.Equal(t, "den@yandex.ru", req.Email)
					require.Equal(t, "Denis", req.Name)
					require.Equal(t, "Марс", req.Address)
					err := bcrypt.CompareHashAndPassword([]byte(req.Password), []byte("12345678"))
					require.NoError(t, err, "password should be a valid bcrypt hash of 12345678")
					return userID, nil
				})
				return mock
			},
		},
		{
			name: "duplicate email error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  fmt.Errorf("email already registered"),
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Set(func(ctx context.Context, req *models.RegisterRequest) (string, error) {
					return "", errors.New("duplicate key value violates unique constraint")
				})
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := tt.userRepositoryMock(mc)
			srv := service.NewService(userRepoMock, "secret")

			newID, err := srv.RegisterUser(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, newID)
		})
	}
}
