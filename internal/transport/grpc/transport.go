package transport

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/ecommerc-go/users/internal/lib/validation"
	"github.com/ecommerc-go/users/internal/service"
	"github.com/ecommerc-go/users/pkg/users"
)

var (
	ErrValidation = errors.New("validation error")
)

type Implementation struct {
	users.UnimplementedUserServiceServer
	service  service.UserService
	validate *validation.Validator
}

// NewImplementation создает новый экземпляр gRPC реализации
func NewImplementation(svc service.UserService) *Implementation {
	return &Implementation{
		service:  svc,
		validate: validation.NewValidator(),
	}
}

// RegisterUser обрабатывает запрос на регистрацию пользователя
func (s *Implementation) RegisterUser(ctx context.Context, req *users.RegisterUserRequest) (*users.RegisterUserResponse, error) {
	logger := slog.With("operation", OpRegisterUser, "email", req.GetEmail())

	// Валидация входных данных
	if err := s.validate.ValidateRegisterUser(req); err != nil {
		logger.Warn("validation error", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrValidation, err)
	}

	logger.Debug("starting user registration")

	ID, err := s.service.RegisterUser(ctx, RegisterUserFromProto(req))
	if err != nil {
		logger.Error("registration failed", "error", err)
		return nil, err
	}

	logger.Debug("registration completed", "user_id", ID)
	return &users.RegisterUserResponse{UserId: ID}, nil
}

// LoginUser обрабатывает запрос на аутентификацию пользователя
func (s *Implementation) LoginUser(ctx context.Context, req *users.LoginUserRequest) (*users.LoginUserResponse, error) {
	logger := slog.With("operation", OpLoginUser, "email", req.GetEmail())

	// Валидация входных данных
	if err := s.validate.ValidateLoginUser(req); err != nil {
		logger.Warn("validation failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrValidation, err)
	}

	logger.Debug("starting user login")
	token, err := s.service.LoginUser(ctx, LoginUserFromProto(req))
	if err != nil {
		logger.Error("login failed", "error", err)
		return nil, err
	}

	logger.Debug("login successful")
	return &users.LoginUserResponse{JwtToken: token}, nil
}

// GetProfile обрабатывает запрос на получение профиля пользователя
func (s *Implementation) GetProfile(ctx context.Context, req *users.GetProfileRequest) (*users.GetProfileResponse, error) {
	logger := slog.With("operation", OpGetProfile, "user_id", req.GetUserId())

	// Валидация входных данных
	if err := s.validate.ValidateUserID(req.GetUserId()); err != nil {
		logger.Warn("validation failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrValidation, err)
	}

	logger.Debug("fetching user profile")
	profile, err := s.service.GetProfile(ctx, req.GetUserId())
	if err != nil {
		logger.Error("failed to get profile", "error", err)
		return nil, err
	}

	logger.Debug("profile retrieved successfully")
	return UserProfileToProto(profile), nil
}

// DeleteProfile обрабатывает запрос на удаление профиля пользователя
func (s *Implementation) DeleteProfile(ctx context.Context, req *users.DeleteProfileRequest) (*users.DeleteProfileResponse, error) {
	logger := slog.With("operation", OpDeleteProfile, "user_id", req.GetUserId())

	// Валидация входных данных
	if err := s.validate.ValidateUserID(req.GetUserId()); err != nil {
		logger.Warn("validation failed", "error", err)
		return &users.DeleteProfileResponse{Success: false}, fmt.Errorf("%w: %v", ErrValidation, err)
	}

	logger.Debug("starting profile deletion")
	err := s.service.DeleteProfile(ctx, req.GetUserId())
	if err != nil {
		logger.Error("failed to delete profile", "error", err)
		return &users.DeleteProfileResponse{Success: false}, err
	}

	logger.Debug("profile deleted successfully")
	return &users.DeleteProfileResponse{Success: true}, nil
}

// UpdateProfile обрабатывает запрос на редактирование  профиля пользователя
func (s *Implementation) UpdateProfile(ctx context.Context, req *users.UpdateProfileRequest) (*users.UpdateProfileResponse, error) {
	logger := slog.With("operation", OpUpdateProfile, "user_id", req.GetUserId())

	// Валидация входных данных
	if err := s.validate.ValidateUpdateProfile(req); err != nil {
		logger.Warn("validation failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrValidation, err)
	}

	err := s.service.UpdateProfile(ctx, UpdateProfileFromProto(req))
	if err != nil {
		return &users.UpdateProfileResponse{
			Success: false,
		}, err
	}

	logger.Debug("profile retrieved successfully")
	return &users.UpdateProfileResponse{
		Success: true,
	}, nil
}
