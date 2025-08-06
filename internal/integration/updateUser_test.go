package integration

import (
	"context"
	"testing"
	"time"

	"github.com/ecommerc-go/users/internal/domain"
	utils_test "github.com/ecommerc-go/users/internal/integration/utils"
	repository "github.com/ecommerc-go/users/internal/repository/postgress"
	"github.com/ecommerc-go/users/internal/service"
	transport "github.com/ecommerc-go/users/internal/transport/grpc"
	"github.com/ecommerc-go/users/pkg/users"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser_Integration(t *testing.T) {

	db := utils_test.SetupTestDB(t)
	user := &domain.RegisterUser{
		Email:    "test@example.com",
		Password: "123456789",
		Name:     "DEN",
		Address:  "Neznayakovo",
	}

	ID := utils_test.CreateTestUser(db, t, user)
	defer utils_test.ClearTestUsers(db, t)

	repo := repository.NewRepository(db)
	svc := service.NewService(repo, "test-secret")
	api := transport.NewImplementation(svc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &users.UpdateProfileRequest{
		UserId:  ID,
		Name:    "Andrey",
		Address: "Andreevskaya",
	}

	resp, err := api.UpdateProfile(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	var name, address string
	err = db.QueryRowContext(ctx, "SELECT name, address FROM users WHERE email = $1", "test@example.com").Scan(&name, &address)
	assert.NoError(t, err)
	assert.Equal(t, req.Name, name)
	assert.Equal(t, req.Address, address)

}
