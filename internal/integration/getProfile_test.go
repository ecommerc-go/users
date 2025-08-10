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
	"golang.org/x/crypto/bcrypt"
)

func TestGetProfile_Integration(t *testing.T) {

	db := utils_test.SetupTestDB(t)
	user := &domain.RegisterUser{
		Email:    "test@example.com",
		Password: "123456789",
		Name:     "DEN",
		Address:  "Neznayakovo",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456789"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)

	ID := utils_test.CreateTestUser(db, t, user)
	defer utils_test.ClearTestUsers(db, t)

	repo := repository.NewRepository(db)
	svc := service.NewService(repo, "test-secret")
	api := transport.NewImplementation(svc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &users.GetProfileRequest{
		UserId: ID,
	}

	resp, err := api.GetProfile(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	var name, email string
	err = db.QueryRowContext(ctx, "SELECT email, name FROM users WHERE id = $1", ID).Scan(&email, &name)
	assert.NoError(t, err)
	assert.Equal(t, resp.Profile.Email, email)
	assert.Equal(t, resp.Profile.Name, name)

}
