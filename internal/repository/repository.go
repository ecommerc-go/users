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
