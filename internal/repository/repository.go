package repository

import (
	"context"
	"errors"
	"fmt"

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

func (r *Repository) CreateUser(ctx context.Context, user *RepoRegisterReq) (string, error) {
	query := `INSERT INTO users (email, password_hash, name, address) 
	          VALUES ($1, $2, $3, $4)
	          RETURNING id`

	var UserId string
	err := r.conn.QueryRowContext(ctx, query,
		user.Email, user.Password, user.Name, user.Address).
		Scan(&UserId)
	fmt.Println(err)
	if err != nil {
		//лог
		return "", errors.New("failed to create user")
	}

	return UserId, nil
}

func (r *Repository) GetUser(ctx context.Context, id string) (*RepoUserProfile, error) {
	query := `SELECT email, name, address 
	          FROM users WHERE id=$1`

	var profile RepoUserProfile
	err := r.conn.QueryRowContext(ctx, query,
		id).
		Scan(&profile.Email, &profile.Name, &profile.Address)
	fmt.Println(err)
	if err != nil {
		// добавить лог
		return nil, errors.New("failed to get user")
	}
	return &profile, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id=$1;`

	_, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		//лог
		return errors.New("user dont delete")
	}
	return nil
}
