package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/ecommerc-go/users/internal/domain"
	srvMock "github.com/ecommerc-go/users/internal/service/mocks"
	transport "github.com/ecommerc-go/users/internal/transport/grpc"
	"github.com/ecommerc-go/users/pkg/users"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestLoginUser(t *testing.T) {
	var (
		ctx      = context.Background()
		mc       = minimock.NewController(t)
		email    = "valid@example.com"
		password = "securePassword123"
		token    = "generated.jwt.token"
	)

	tests := []struct {
		name      string
		req       *users.LoginUserRequest
		mockSetup func(*srvMock.UserServiceMock)
		want      *users.LoginUserResponse
		wantErr   bool
		errorText string
	}{
		{
			name: "successful login",
			req: &users.LoginUserRequest{
				Email:    email,
				Password: password,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.LoginUserMock.Expect(ctx, &domain.LoginUser{
					Email:    email,
					Password: password,
				}).Return(token, nil)
			},
			want: &users.LoginUserResponse{
				JwtToken: token,
			},
			wantErr: false,
		},
		{
			name: "invalid email format",
			req: &users.LoginUserRequest{
				Email:    "invalid-email",
				Password: password,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {},
			wantErr:   true,
			errorText: "validation error",
		},
		{
			name: "empty password",
			req: &users.LoginUserRequest{
				Email:    email,
				Password: "",
			},
			mockSetup: func(m *srvMock.UserServiceMock) {},
			wantErr:   true,
			errorText: "validation error",
		},
		{
			name: "invalid credentials",
			req: &users.LoginUserRequest{
				Email:    email,
				Password: password,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.LoginUserMock.Return("", errors.New("invalid credentials"))
			},
			wantErr:   true,
			errorText: "invalid credentials",
		},
		{
			name: "service error",
			req: &users.LoginUserRequest{
				Email:    email,
				Password: password,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.LoginUserMock.Return("", errors.New("login failed"))
			},
			wantErr:   true,
			errorText: "login failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSrv := srvMock.NewUserServiceMock(mc)
			tt.mockSetup(mockSrv)

			tran := transport.NewImplementation(mockSrv)

			resp, err := tran.LoginUser(ctx, tt.req)

			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorText)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, resp)
			}
		})
	}
}
