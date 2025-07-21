package test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ecommerc-go/users/internal/models"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestUpdateUserl(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx context.Context
		req *models.UpdateProfileRequest
	}

	tests := []struct {
		name   string
		args   args
		want   int64
		err    error
		mockDB func(mock sqlmock.Sqlmock)
	}{
		{
			name: "update success",
			args: args{
				ctx: ctx,
				req: &models.UpdateProfileRequest{
					ID:      "f47ac10b-58cc-4372-a567-0e02b2c3d479",
					Name:    "DEN",
					Address: "NEZNAYKA",
				},
			},
			want: 1,
			err:  nil,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE users`).
					WithArgs("DEN", "NEZNAYKA", "f47ac10b-58cc-4372-a567-0e02b2c3d479").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "update rows 0",
			args: args{
				ctx: ctx,
				req: &models.UpdateProfileRequest{
					ID:      "f47ac10b-58cc-4372-a567-0e02b2c3d479",
					Name:    "DEN",
					Address: "NEZNAYKA",
				},
			},
			want: 0,
			err:  nil,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE users`).
					WithArgs("DEN", "NEZNAYKA", "f47ac10b-58cc-4372-a567-0e02b2c3d479").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
		{
			name: "update user error",
			args: args{
				ctx: ctx,
				req: &models.UpdateProfileRequest{
					ID:      "f47ac10b-58cc-4372-a567-0e02b2c3d479",
					Name:    "DEN",
					Address: "NEZNAYKA",
				},
			},
			want: 0,
			err:  errors.New("error database"),
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE users`).
					WithArgs("DEN", "NEZNAYKA", "f47ac10b-58cc-4372-a567-0e02b2c3d479").
					WillReturnError(errors.New("error database"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)

			tt.mockDB(mock)
			sqlxDB := sqlx.NewDb(db, "sqlmock")
			repo := repository.NewRepository(sqlxDB)

			profile, err := repo.UpdateUser(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, profile)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
