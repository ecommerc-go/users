package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ecommerc-go/users/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	conn *sqlx.DB
}

var (
	ErrAlreadyExist = errors.New("username already exists")
)

func NewRepository(conn *sqlx.DB) *Repository {

	return &Repository{conn: conn}
}

func (r *Repository) CreateUser(ctx context.Context, user *service.RegisterUserRequest) (*service.RegisterUserResponse, error) {
	query := `INSERT INTO users (email, password_hash, name, address) 
	          VALUES ($1, $2, $3, $4)
	          RETURNING id`

	var createdUser string
	err := r.conn.QueryRowContext(ctx, query,
		user.Email, user.Password, user.Name, user.Address).
		Scan(&createdUser)
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("failed to create user")
	}
	fmt.Println(createdUser)
	return &service.RegisterUserResponse{
		Id: string(createdUser),
	}, nil
}

func (r *Repository) GetUser(ctx context.Context, id string) (*service.GetProfileResponse, error) {
	query := `SELECT email, name, address 
	          FROM users WHERE id=$1`

	var profile service.UserProfile
	err := r.conn.QueryRowContext(ctx, query,
		id).
		Scan(&profile.Email, &profile.Name, &profile.Address)
	fmt.Println(err)
	fmt.Println()
	if err != nil {
		return nil, errors.New("failed to create user")
	}
	return &service.GetProfileResponse{
		Profile: profile,
	}, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id=$1;`

	_, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return errors.New("user dont delete")
	}
	return nil
}
