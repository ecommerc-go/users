package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/ecommerc-go/users/internal/domain"
	repoMocks "github.com/ecommerc-go/users/internal/repository/postgress/mocks"
	"github.com/ecommerc-go/users/internal/service"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGetProfile(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		email   = "den@yandex.ru"
		address = "NEZNAYKA"
		name    = "Den"
		userID  = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	)

	tests := []struct {
		name        string
		args        args
		mockSetup   func(*repoMocks.UserRepositoryMock)
		want        *domain.UserProfile
		wantErr     bool
		errContains string
	}{
		{
			name: "get profile success",
			args: args{ctx, userID},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.GetUserMock.Expect(ctx, userID).Return(
					&domain.UserProfile{
						ID:      userID,
						Email:   email,
						Address: address,
						Name:    name,
					}, nil)
			},
			want: &domain.UserProfile{
				ID:      userID,
				Email:   email,
				Address: address,
				Name:    name,
			},
			wantErr: false,
		},
		{
			name: "profile not found",
			args: args{ctx, userID},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.GetUserMock.Expect(ctx, userID).Return(
					nil, errors.New("profile not found"))
			},
			want:        nil,
			wantErr:     true,
			errContains: "profile not found",
		},
		{
			name: "repository error",
			args: args{ctx, userID},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.GetUserMock.Expect(ctx, userID).Return(
					nil, errors.New("db error"))
			},
			want:        nil,
			wantErr:     true,
			errContains: "failed to get profile",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repoMocks.NewUserRepositoryMock(mc)
			tt.mockSetup(mockRepo)

			srv := service.NewService(mockRepo, "secret")

			got, err := srv.GetProfile(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errContains)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
