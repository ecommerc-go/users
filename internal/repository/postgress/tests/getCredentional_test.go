package test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestGetCredentional(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx   context.Context
		email string
	}

	tests := []struct {
		name   string
		args   args
		want   *repository.Creds
		err    error
		mockDB func(mock sqlmock.Sqlmock)
	}{
		{
			name: "get creds success",
			args: args{
				ctx:   ctx,
				email: "den@yandex.ru",
			},
			want: &repository.Creds{
				Login:    "den@yandex.ru",
				Password: "123456789",
				ID:       "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			},
			err: nil,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT id ,email, password_hash FROM users`).
					WithArgs("den@yandex.ru").
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password_hash"}).
						AddRow("f47ac10b-58cc-4372-a567-0e02b2c3d479", "den@yandex.ru", "123456789"))
			},
		},
		{
			name: "get creds error",
			args: args{
				ctx:   ctx,
				email: "den@yandex.ru",
			},
			want: &repository.Creds{
				Login:    "den@yandex.ru",
				Password: "123456789",
				ID:       "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			},
			err: errors.New("failed to getCreds in BD"),
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT id ,email, password_hash FROM users`).
					WithArgs("den@yandex.ru").
					WillReturnError(errors.New("failed to getCreds in BD"))
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

			profile, err := repo.GetCredentials(tt.args.ctx, tt.args.email)
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
