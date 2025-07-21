package tests

import (
	"context"
	"errors"
	"testing"

	srvMock "github.com/ecommerc-go/users/internal/service/mocks"
	transport "github.com/ecommerc-go/users/internal/transport/grpc"
	"github.com/ecommerc-go/users/pkg/users"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDeleteProfile(t *testing.T) {
	var (
		ctx    = context.Background()
		mc     = minimock.NewController(t)
		userID = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	)

	tests := []struct {
		name      string
		req       *users.DeleteProfileRequest
		mockSetup func(*srvMock.UserServiceMock)
		want      *users.DeleteProfileResponse
		wantErr   bool
		errorText string
	}{
		{
			name: "successful deletion",
			req: &users.DeleteProfileRequest{
				UserId: userID,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.DeleteProfileMock.Expect(ctx, userID).Return(nil)
			},
			want: &users.DeleteProfileResponse{
				Success: true,
			},
			wantErr: false,
		},
		{
			name: "invalid ID",
			req: &users.DeleteProfileRequest{
				UserId: "1",
			},
			mockSetup: func(m *srvMock.UserServiceMock) {},
			want: &users.DeleteProfileResponse{
				Success: false,
			},
			wantErr:   true,
			errorText: "validation error",
		},
		{
			name: "service error",
			req: &users.DeleteProfileRequest{
				UserId: userID,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.DeleteProfileMock.Return(errors.New("service error"))
			},
			want: &users.DeleteProfileResponse{
				Success: false,
			},
			wantErr:   true,
			errorText: "service error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSrv := srvMock.NewUserServiceMock(mc)
			tt.mockSetup(mockSrv)

			tran := transport.NewImplementation(mockSrv)

			resp, err := tran.DeleteProfile(ctx, tt.req)

			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorText)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, resp)
		})
	}
}
