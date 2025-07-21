package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/ecommerc-go/users/internal/models"
	srvMock "github.com/ecommerc-go/users/internal/service/mocks"
	transport "github.com/ecommerc-go/users/internal/transport/grpc"
	"github.com/ecommerc-go/users/pkg/users"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {

	var (
		ctx      = context.Background()
		mc       = minimock.NewController(t)
		userID   = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
		address  = "NEZNAYKA"
		password = "123456890"
		email    = "den@yandex.ru"
		name     = "Den"
	)

	tests := []struct {
		ctx       context.Context
		name      string
		req       *users.RegisterUserRequest
		want      *users.RegisterUserResponse
		mockSetup func(*srvMock.UserServiceMock)
		wantErr   bool
		errorText string
	}{
		{
			ctx:  ctx,
			name: "register user succes",
			req: &users.RegisterUserRequest{
				Name:     name,
				Email:    email,
				Password: password,
				Address:  address,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.RegisterUserMock.Expect(ctx, &models.RegisterRequest{
					Email:    email,
					Password: password,
					Name:     name,
					Address:  address,
				}).Return(userID, nil)
			},
			want: &users.RegisterUserResponse{
				UserId: userID,
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			req: &users.RegisterUserRequest{
				Name:     name,
				Email:    "qwerty",
				Password: password,
				Address:  address,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {

			},
			want:      nil,
			wantErr:   true,
			errorText: "validation error",
		},
		{
			name: "invalid password",
			req: &users.RegisterUserRequest{
				Name:     name,
				Email:    email,
				Password: "12345",
				Address:  address,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {

			},
			want:      nil,
			wantErr:   true,
			errorText: "validation error",
		},
		{
			name: "registration failed",
			ctx:  ctx,
			req: &users.RegisterUserRequest{
				Name:     name,
				Email:    email,
				Password: password,
				Address:  address,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.RegisterUserMock.Expect(ctx, &models.RegisterRequest{
					Email:    email,
					Password: password,
					Name:     name,
					Address:  address,
				}).Return("", errors.New("registration failed"))
			},
			want:      nil,
			wantErr:   true,
			errorText: "registration failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSrv := srvMock.NewUserServiceMock(mc)
			tt.mockSetup(mockSrv)

			tran := transport.NewImplementation(mockSrv)

			ID, err := tran.RegisterUser(tt.ctx, tt.req)

			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorText)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, ID)
		})
	}
}
