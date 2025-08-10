package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/ecommerc-go/users/internal/domain"
	"github.com/ecommerc-go/users/internal/lib/jwt"

	repoMocks "github.com/ecommerc-go/users/internal/repository/postgress/mocks"
	"github.com/ecommerc-go/users/internal/service"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginUser(t *testing.T) {

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
	jwtSrv := jwt.NewJWTService("secret")
	jwtToken := jwtSrv.CreateToken(userID)

	tests := []struct {
		name      string
		args      args
		want      string
		err       error
		mockSetup func(*repoMocks.UserRepositoryMock)
	}{
		{
			name: "login success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: jwtToken,
			err:  nil,
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.GetCredentialsMock.Expect(ctx, req.Email).Return(&domain.Creds{
					ID:       userID,
					Login:    email,
					Password: string(hashedPassword),
				}, nil)
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
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.GetCredentialsMock.Expect(ctx, req.Email).Return(nil, sql.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repoMocks.NewUserRepositoryMock(mc)
			tt.mockSetup(mockRepo)
			srv := service.NewService(mockRepo, "secret")

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
