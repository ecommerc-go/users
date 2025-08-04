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

func TestGetProfile(t *testing.T) {
	var (
		ctx    = context.Background()
		mc     = minimock.NewController(t)
		userID = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	)

	testProfile := &models.UserProfile{
		ID:         userID,
		Email:      "test@example.com",
		Name:       "Test User",
		Address:    "123 Main St",
		Created_at: "2023-01-01T00:00:00Z",
	}

	tests := []struct {
		name      string
		req       *users.GetProfileRequest
		mockSetup func(*srvMock.UserServiceMock)
		want      *users.GetProfileResponse
		wantErr   bool
		errorText string
	}{
		{
			name: "successful get profile",
			req: &users.GetProfileRequest{
				UserId: userID,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.GetProfileMock.Expect(ctx, userID).Return(testProfile, nil)
			},
			want: &users.GetProfileResponse{
				Profile: &users.UserProfile{
					UserId:    testProfile.ID,
					Email:     testProfile.Email,
					Name:      testProfile.Name,
					Address:   testProfile.Address,
					CreatedAt: testProfile.Created_at,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid ID",
			req: &users.GetProfileRequest{
				UserId: "1",
			},
			mockSetup: func(m *srvMock.UserServiceMock) {},
			want:      nil,
			wantErr:   true,
			errorText: "validation error",
		},
		{
			name: "service error",
			req: &users.GetProfileRequest{
				UserId: userID,
			},
			mockSetup: func(m *srvMock.UserServiceMock) {
				m.GetProfileMock.Return(nil, errors.New("service error"))
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

			resp, err := tran.GetProfile(ctx, tt.req)

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
