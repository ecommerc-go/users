package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/ecommerc-go/users/internal/domain"
	"github.com/ecommerc-go/users/internal/lib/jwt"

	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	repoMocks "github.com/ecommerc-go/users/internal/repository/postgress/mocks"
	"github.com/ecommerc-go/users/internal/service"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginUser(t *testing.T) {
	type UserRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *domain.LoginUser
	}

	var (
		ctx      = context.Background()
		mc       = minimock.NewController(t)
		userID   = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
		email    = "den@yandex.ru"
		password = "12345678"
		req      = &domain.LoginUser{
			Email:    email,
			Password: password,
		}
		hashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	)

	tests := []struct {
		name               string
		args               args
		want               string
		err                error
		userRepositoryMock UserRepositoryMockFunc
	}{
		{
			name: "login success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: jwt.CreateJWTToken(userID, "secret"),
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetCredentialsMock.Set(func(ctx context.Context, email string) (*domain.Creds, error) {
					require.Equal(t, "den@yandex.ru", email)
					return &domain.Creds{
						ID:       userID,
						Login:    email,
						Password: string(hashedPassword),
					}, nil
				})
				return mock
			},
		},
		{
			name: "user not found",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  fmt.Errorf("email not found"),
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetCredentialsMock.Set(func(ctx context.Context, email string) (*domain.Creds, error) {
					return nil, sql.ErrNoRows
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

			token, err := srv.LoginUser(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, token)
				if tt.want != "" {
					require.Equal(t, tt.want, token)
				}
			}
		})
	}
}
