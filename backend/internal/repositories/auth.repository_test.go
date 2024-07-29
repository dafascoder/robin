package repositories

import (
	model "backend/internal/models"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthRepository_CreateAccount(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create mock queries
	mockQueries := model.New(db)

	// Create AuthRepository with mock database and queries
	authRepo := &AuthRepository{
		db:      db,
		queries: mockQueries,
	}

	// Define the input parameters
	params := model.CreateAccountParams{
		Email:    "test@example.com",
		Password: "password",
	}

	// Define the expected result
	expectedAccount := model.CreateAccountRow{
		Email: "test@example.com",
	}

	// Set up the mock expectations
	mock.ExpectQuery("INSERT INTO account").
		WithArgs(params.Email, params.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).AddRow(expectedAccount.ID, expectedAccount.Email))

	// Call the CreateAccount method
	account, err := authRepo.CreateAccount(context.Background(), params)

	// Verify the results
	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, account)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
