package utils_test

import (
	"context"
	"testing"

	"github.com/ecommerc-go/users/internal/domain"
	"github.com/jmoiron/sqlx"
)

func CreateTestUser(db *sqlx.DB, t *testing.T, user *domain.RegisterUser) string {
	query := `INSERT INTO users (email, password_hash, name, address) 
	          VALUES ($1, $2, $3, $4) 
	          RETURNING id`
	var id string
	err := db.QueryRowContext(context.Background(), query,
		user.Email, user.Password, user.Name, user.Address).Scan(&id)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	return id
}

func ClearTestUsers(db *sqlx.DB, t *testing.T) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("Failed to clear users table: %v", err)
	}
}
