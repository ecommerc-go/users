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

func TestUpdateProfile(t *testing.T) {
	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		userID  = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
		name    = "Updated Name"
		address = "Updated Address"
	)

	tests := []struct {
		name        string
		req         *users.UpdateProfileRequest
		mockSetup   func(*srvMock.UserServiceMock)
		want        *users.UpdateProfileResponse
		wantErr     bool
		errorText   string
		expectError error
	}{
		{
			name: "successful update",
			req: &users.UpdateProfileRequest{
				UserId:  userID,
				Name:    name,
				Address: address,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.UpdateProfileMock.Expect(ctx, &domain.UpdateProfile{
					ID:      userID,
					Name:    name,
					Address: address,
				}).Return(nil)
			},
			want: &users.UpdateProfileResponse{
				Success: true,
			},
			wantErr: false,
		},
		{
			name: "validation error - empty user ID",
			req: &users.UpdateProfileRequest{
				UserId:  "",
				Name:    name,
				Address: address,
			},
			mockSetup:   func(m *srvMock.UserServiceMock) {},
			want:        nil,
			wantErr:     true,
			errorText:   "validation error",
			expectError: transport.ErrValidation,
		},
		{
			name: "service error",
			req: &users.UpdateProfileRequest{
				UserId:  userID,
				Name:    name,
				Address: address,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.UpdateProfileMock.Return(errors.New("service error"))
			},
			want: &users.UpdateProfileResponse{
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

			impl := transport.NewImplementation(mockSrv)
			resp, err := impl.UpdateProfile(ctx, tt.req)

			if tt.wantErr {
				require.Error(t, err)
				if tt.expectError != nil {
					require.ErrorIs(t, err, tt.expectError)
				}
				if tt.errorText != "" {
					require.Contains(t, err.Error(), tt.errorText)
				}
			} else {
				require.NoError(t, err)
			}

			if tt.want != nil {
				require.Equal(t, tt.want.Success, resp.Success)
			}
		})
	}
}
