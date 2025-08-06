package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ecommerc-go/users/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	conn *sqlx.DB
}

// NewRepository создает новый экземпляр Repository с подключением к БД
func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{conn: conn}
}

// CreateUser создает нового пользователя в базе данных
func (r *Repository) CreateUser(ctx context.Context, user *domain.RegisterUser) (string, error) {
	logger := slog.With(
		"op", OpCreateUser,
		"email", user.Email,
	)

	query := `INSERT INTO users (email, password_hash, name, address) 
	          VALUES ($1, $2, $3, $4)
	          RETURNING id`

	var userID string

	err := r.conn.QueryRowContext(ctx, query,
		user.Email, user.Password, user.Name, user.Address).
		Scan(&userID)

	if err != nil {
		logger.Error("failed to create user", "error", err)
		return "", err
	}

	return userID, nil
}

// GetUser получает профиль пользователя по его ID
func (r *Repository) GetUser(ctx context.Context, id string) (*domain.UserProfile, error) {
	logger := slog.With(
		"op", OpGetUser,
		"user_id", id,
	)

	query := `SELECT email, name, address 
	          FROM users WHERE id=$1`

	var profile domain.UserProfile
	err := r.conn.QueryRowContext(ctx, query,
		id).
		Scan(&profile.Email, &profile.Name, &profile.Address)

	if err != nil {
		logger.Error("failed to get user in BD", "error", err)
		return nil, err
	}
	return &profile, nil
}

// DeleteUser удаляет пользователя по его ID
func (r *Repository) DeleteUser(ctx context.Context, ID string) error {
	logger := slog.With(
		"op", OpDeleteUser,
		"user_id", ID,
	)
	query := `DELETE FROM users WHERE id=$1;`

	_, err := r.conn.ExecContext(ctx, query, ID)
	if err != nil {
		logger.Error("failed to delete user in BD", "error", err)
		return err
	}
	return nil
}

// GetCredentials получает учетные данные пользователя по email
func (r *Repository) GetCredentials(ctx context.Context, email string) (*domain.Creds, error) {
	logger := slog.With(
		"op", OpGetCredentials,
		"email", email,
	)
	query := `SELECT id ,email, password_hash FROM users WHERE email=$1`
	var (
		ID, DBemail, DBPassword string
	)

	err := r.conn.QueryRowContext(ctx, query, email).Scan(&ID, &DBemail, &DBPassword)
	if err != nil {
		logger.Error("failed to getCreds in BD", "error", err)
		return nil, err
	}

	return &domain.Creds{
		Login:    DBemail,
		Password: DBPassword,
		ID:       ID,
	}, nil
}

// UpdateUser обновляет все данные пользователя в базе данных
func (r *Repository) UpdateUser(ctx context.Context, user *domain.UpdateProfile) (int64, error) {
	logger := slog.With(
		"op", OpUpdateUser,
		"user_id", user.ID,
	)

	query := `UPDATE users 
              SET name = $1, address = $2
              WHERE id = $3`

	logger.Debug("update query",
		"query", query,
		"name", user.Name,
		"address", user.Address)

	result, err := r.conn.ExecContext(ctx, query,
		user.Name,
		user.Address,
		user.ID)

	if err != nil {
		logger.Error("failed to execute update query", "error", err)
		return 0, fmt.Errorf("failed to update user in BD: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("failed to get affected rows", "error", err)
		return 0, fmt.Errorf("ailed to get affected rows: %w", err)
	}

	return rowsAffected, nil
}
