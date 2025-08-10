package test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ecommerc-go/users/internal/domain"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx  context.Context
		user *domain.RegisterUser
	}

	tests := []struct {
		name   string
		args   args
		want   string
		err    error
		mockDB func(mock sqlmock.Sqlmock)
	}{
		{
			name: "create user succes",
			args: args{
				ctx: ctx,
				user: &domain.RegisterUser{
					Email:    "den@yandex.ru",
					Password: "hashedpassword",
					Name:     "DEN",
					Address:  "NEZNAYKA",
				},
			},
			want: "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			err:  nil,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO users`).
					WithArgs("den@yandex.ru", "hashedpassword", "DEN", "NEZNAYKA").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("f47ac10b-58cc-4372-a567-0e02b2c3d479"))
			},
		},
		{
			name: "create user error",
			args: args{
				ctx: ctx,
				user: &domain.RegisterUser{
					Email:    "den@yandex.ru",
					Password: "hashedpassword",
					Name:     "DEN",
					Address:  "NEZNAYKA",
				},
			},
			want: "",
			err:  errors.New("ошибка базы данных"),
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO users`).
					WithArgs("den@yandex.ru", "hashedpassword", "DEN", "NEZNAYKA").
					WillReturnError(errors.New("ошибка базы данных"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer func() {
				if err := db.Close(); err != nil {
					fmt.Printf("failed to close database: %v", err)
				}
			}()

			tt.mockDB(mock)
			sqlxDB := sqlx.NewDb(db, "sqlmock")
			repo := repository.NewRepository(sqlxDB)

			userID, err := repo.CreateUser(tt.args.ctx, tt.args.user)
			if tt.err != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, userID)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
