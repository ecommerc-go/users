package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	utils_test "github.com/ecommerc-go/users/internal/integration/utils"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"github.com/ecommerc-go/users/internal/service"
	transport "github.com/ecommerc-go/users/internal/transport/grpc"
	"github.com/ecommerc-go/users/pkg/users"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser_Integration(t *testing.T) {

	db := utils_test.SetupTestDB(t)
	defer utils_test.ClearTestUsers(db, t)

	repo := repository.NewRepository(db)
	svc := service.NewService(repo, "test-secret")
	api := transport.NewImplementation(svc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &users.RegisterUserRequest{
		Address:  "Neznaykino",
		Email:    "test@example.com",
		Password: "12345678",
		Name:     "Test User",
	}

	id, err := api.RegisterUser(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	var name, password string

	err = db.QueryRowContext(ctx, "SELECT name FROM users WHERE email = $1", "test@example.com").Scan(&name)
	fmt.Println(name)
	fmt.Println(password)
	assert.NoError(t, err)
	assert.Equal(t, "Test User", name)
}
