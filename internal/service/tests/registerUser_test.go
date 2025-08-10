package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/ecommerc-go/users/internal/domain"
	repoMocks "github.com/ecommerc-go/users/internal/repository/postgress/mocks"
	"github.com/ecommerc-go/users/internal/service"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {

	type args struct {
		ctx context.Context
		req *domain.RegisterUser
	}

	var (
		ctx    = context.Background()
		mc     = minimock.NewController(t)
		userID = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
		req    = &domain.RegisterUser{
			Email:    "den@yandex.ru",
			Password: "12345678",
			Name:     "Denis",
			Address:  "Марс",
		}
	)

	tests := []struct {
		name      string
		args      args
		want      string
		err       error
		mockSetup func(*repoMocks.UserRepositoryMock)
	}{
		{
			name: "register success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: userID,
			err:  nil,
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.CreateUserMock.Expect(ctx, req).Return(userID, nil)
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
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.CreateUserMock.Expect(ctx, req).Return("", errors.New("duplicate key value violates unique constraint"))
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repoMocks.NewUserRepositoryMock(mc)
			tt.mockSetup(mockRepo)
			srv := service.NewService(mockRepo, "secret")

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
