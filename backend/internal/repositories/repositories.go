package repositories

import (
	model "backend/internal/models"
	"context"
	"github.com/google/uuid"
)

// interface below

type AuthRepositoryInterface interface {
	CreateAccount(ctx context.Context, params model.CreateAccountParams) (model.CreateAccountRow, error)
	GetAccountByEmail(ctx context.Context, email string) (model.Account, error)
	GetAccountByID(ctx context.Context, id uuid.UUID) (model.Account, error)
	VerifyAccount(ctx context.Context, token string) error
}

// constructors below

func NewAuthRepository(queries *model.Queries) *AuthRepository {
	return &AuthRepository{
		queries: queries,
	}
}
