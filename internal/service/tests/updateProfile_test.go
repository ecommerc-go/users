package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/ecommerc-go/users/internal/models"
	repoMocks "github.com/ecommerc-go/users/internal/repository/postgress/mocks"
	"github.com/ecommerc-go/users/internal/service"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestUpdateProfile(t *testing.T) {
	type args struct {
		ctx context.Context
		req *models.UpdateProfileRequest
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		userID  = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
		address = "NEZNAYKA"
		name    = "Den"
		req     = &models.UpdateProfileRequest{
			ID:      userID,
			Name:    name,
			Address: address,
		}
	)

	tests := []struct {
		name      string
		args      args
		mockSetup func(*repoMocks.UserRepositoryMock)
		wantErr   bool
		errorText string
	}{
		{
			name: "success case",
			args: args{ctx, req},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.UpdateUserMock.Expect(ctx, req).Return(1, nil)
			},
			wantErr: false,
		},
		{
			name: "repository error",
			args: args{ctx, req},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.UpdateUserMock.Expect(ctx, req).Return(
					0, errors.New("update failed"))
			},
			wantErr:   true,
			errorText: "update failed",
		},
		{
			name: "user not found",
			args: args{ctx, req},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.UpdateUserMock.Expect(ctx, req).Return(
					0, nil)
			},
			wantErr:   true,
			errorText: "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repoMocks.NewUserRepositoryMock(mc)
			tt.mockSetup(mockRepo)

			srv := service.NewService(mockRepo, "secret")

			err := srv.UpdateProfile(tt.args.ctx, tt.args.req)
			fmt.Println(err)
			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorText)
			} else {
				require.NoError(t, err)
			}

		})
	}
}
