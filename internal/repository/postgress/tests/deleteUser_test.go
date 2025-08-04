package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestGDeleteUser(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx context.Context
		id  string
	}

	tests := []struct {
		name   string
		args   args
		err    error
		mockDB func(mock sqlmock.Sqlmock)
	}{
		{
			name: "delete user success",
			args: args{
				ctx: ctx,
				id:  "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			},
			err: nil,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM users`).
					WithArgs("f47ac10b-58cc-4372-a567-0e02b2c3d479").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "delete user error",
			args: args{
				ctx: ctx,
				id:  "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			},
			err: sql.ErrNoRows,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM users`).
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

			err = repo.DeleteUser(tt.args.ctx, tt.args.id)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
