package tests

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	repoMocks "github.com/ecommerc-go/users/internal/repository/postgress/mocks"
	"github.com/ecommerc-go/users/internal/service"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDeleteProfile(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}

	var (
		ctx    = context.Background()
		mc     = minimock.NewController(t)
		userId = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	)

	tests := []struct {
		name      string
		args      args
		mockSetup func(*repoMocks.UserRepositoryMock)
		wantErr   bool
		errText   string
	}{
		{
			name: "success case",
			args: args{ctx, userId},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.DeleteUserMock.Expect(ctx, userId).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "user not found",
			args: args{ctx, userId},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.DeleteUserMock.Expect(ctx, userId).Return(
					sql.ErrNoRows)
			},
			wantErr: true,
			errText: "user not found",
		},
		{
			name: "repository error",
			args: args{ctx, userId},
			mockSetup: func(m *repoMocks.UserRepositoryMock) {
				m.DeleteUserMock.Expect(ctx, userId).Return(
					errors.New("failed to delete profile"))
			},
			wantErr: true,
			errText: "failed to delete profile",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repoMocks.NewUserRepositoryMock(mc)
			tt.mockSetup(mockRepo)

			srv := service.NewService(mockRepo, "secret")

			err := srv.DeleteProfile(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errText)
			} else {
				require.NoError(t, err)
			}

		})
	}
}
