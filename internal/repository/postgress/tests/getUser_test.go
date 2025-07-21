package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ecommerc-go/users/internal/models"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx context.Context
		id  string
	}

	tests := []struct {
		name   string
		args   args
		want   *models.UserProfile
		err    error
		mockDB func(mock sqlmock.Sqlmock)
	}{
		{
			name: "get user success",
			args: args{
				ctx: ctx,
				id:  "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			},
			want: &models.UserProfile{
				Email:   "den@yandex.ru",
				Name:    "DEN",
				Address: "NEZNAYKA",
			},
			err: nil,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT email, name, address FROM users`).
					WithArgs("f47ac10b-58cc-4372-a567-0e02b2c3d479").
					WillReturnRows(sqlmock.NewRows([]string{"email", "name", "address"}).
						AddRow("den@yandex.ru", "DEN", "NEZNAYKA"))
			},
		},
		{
			name: "user not found",
			args: args{
				ctx: ctx,
				id:  "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			},
			want: nil,
			err:  sql.ErrNoRows,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT email, name, address FROM users`).
					WithArgs("f47ac10b-58cc-4372-a567-0e02b2c3d479").
					WillReturnError(sql.ErrNoRows)
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

			profile, err := repo.GetUser(tt.args.ctx, tt.args.id)
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
