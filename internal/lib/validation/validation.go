package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ecommerc-go/users/pkg/users"
	"github.com/go-playground/validator"
)

// Validator обертка над validator.Validate
type Validator struct {
	validate *validator.Validate
}

// NewValidator создает новый экземпляр валидатора
func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// RegisterUserRequest структура для валидации запроса на регистрацию
type RegisterUserRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	Name     string `validate:"required"`
	Address  string `validate:"required"`
}

// LoginUserRequest структура для валидации запроса на аутентификацию
type LoginUserRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

// UpdateProfileRequest структура для валидации запроса на обновление профиля
type UpdateProfileRequest struct {
	UserID  string `validate:"required"`
	Name    string `validate:"required"`
	Address string `validate:"required"`
}

// ValidateRegisterUser валидирует запрос на регистрацию
func (v *Validator) ValidateRegisterUser(req *users.RegisterUserRequest) error {

	res := &RegisterUserRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Address:  req.GetAddress(),
	}

	return v.validate.Struct(res)
}

// ValidateLoginUser валидирует запрос на аутентификацию
func (v *Validator) ValidateLoginUser(req *users.LoginUserRequest) error {

	res := &LoginUserRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	return v.validate.Struct(res)
}

// ValidateUserId валидирует user_id
func (v *Validator) ValidateUserID(userId string) error {
	cleanID := strings.Trim(userId, `"`)

	if strings.TrimSpace(cleanID) == "" {
		return errors.New("user ID cannot be empty")
	}

	if cleanID != "" {
		if err := v.validate.Var(cleanID, "uuid"); err != nil {
			return fmt.Errorf("invalid UUID format: %v", err)
		}
	}

	return nil
}

// ValidateUpdateProfile валидирует запрос на обновление профиля
func (v *Validator) ValidateUpdateProfile(req *users.UpdateProfileRequest) error {
	res := &UpdateProfileRequest{
		UserID:  req.GetUserId(),
		Name:    req.GetName(),
		Address: req.GetAddress(),
	}

	return v.validate.Struct(res)
}
