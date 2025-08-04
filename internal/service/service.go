package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ecommerc-go/users/internal/lib/jwt"
	"github.com/ecommerc-go/users/internal/models"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotFound = errors.New("user not found")

type Service struct {
	db         repository.UserRepository
	JWT_SECRET string
}

// NewService создает новый экземпляр Service
func NewService(db repository.UserRepository, JWTSecret string) *Service {
	return &Service{
		db:         db,
		JWT_SECRET: JWTSecret,
	}
}

// RegisterUser регистрирует нового пользователя
func (r *Service) RegisterUser(ctx context.Context, req *models.RegisterRequest) (string, error) {
	logger := slog.With(
		"op", OpRegisterUser,
		"email", req.Email,
	)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to hash password", "error", err)
		return "", errors.New("registration failed")
	}
	req.Password = string(hashedPassword)

	id, err := r.db.CreateUser(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			logger.Warn("user already exists")
			return "", fmt.Errorf("email already registered: : %w", err)
		}

		logger.Error("failed to create user", "error", err)
		return "", fmt.Errorf("registration failed: %w", err)
	}

	logger.Info("user registered successfully")

	return id, nil
}

// LoginUser выполняет аутентификацию пользователя
func (r *Service) LoginUser(ctx context.Context, req *models.LoginUserRequest) (string, error) {
	logger := slog.With(
		"op", OpLoginUser,
		"email", req.Email,
	)

	// 1. Проверка существования email
	creds, err := r.db.GetCredentials(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Warn("email not found")
			return "", fmt.Errorf("email not found:%w", err)
		}
		logger.Error("failed to get email", "err", err)
		return "", fmt.Errorf("failed to get email: %w", err)
	}

	// 2. Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(creds.Password), []byte(req.Password)); err != nil {
		logger.Error("Invalid password attempt for email")
		return "", errors.New("invalid credentials")
	}

	// 3. Генерация токена
	jwtToken := jwt.CreateJWTToken(creds.ID, r.JWT_SECRET)

	return jwtToken, nil
}

// GetProfile получает профиль пользователя по ID
func (r *Service) GetProfile(ctx context.Context, id string) (*models.UserProfile, error) {
	logger := slog.With(
		"op", OpGetProfile,
		"user_id", id,
	)

	data, err := r.db.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Warn("profile not found")
			return nil, fmt.Errorf("profile not found:%w", err)
		}
		logger.Error("failed to get profile", "error", err)
		return nil, fmt.Errorf("failed to get profile:%w", err)
	}
	return data, nil
}

// DeleteProfile удаляет профиль пользователя по ID
func (r *Service) DeleteProfile(ctx context.Context, id string) error {
	logger := slog.With(
		"op", OpDeleteProfile,
		"user_id", id,
	)

	err := r.db.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Warn("user not found")
			return fmt.Errorf("user not found:%w", err)
		}
		logger.Error("failed to delete profile", "error", err)
		return fmt.Errorf("failed to delete profile:%w", err)
	}
	return nil
}

// UpdateProfile редактирует профиль ( name,address)
func (r *Service) UpdateProfile(ctx context.Context, user *models.UpdateProfileRequest) error {
	logger := slog.With(
		"op", OpUpdateProfile,
		"user_id", user.ID,
	)

	rowsAffected, err := r.db.UpdateUser(ctx, user)
	if err != nil {
		logger.Error("update failed", "error", err)
		return fmt.Errorf("update failed: %w", err)
	}

	if rowsAffected == 0 {
		logger.Warn("user not found")
		return ErrUserNotFound
	}

	logger.Info("user updated successfully", "rows_affected", rowsAffected)
	return nil
}
