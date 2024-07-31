package repositories

import (
	model "backend/internal/models"
	"context"
)

// interface below

type AuthRepositoryInterface interface {
	CreateAccount(ctx context.Context, params model.CreateAccountParams) (model.CreateAccountRow, error)
	GetAccountByEmail(ctx context.Context, email string) (model.Account, error)
	GetAccountByID(ctx context.Context) (model.Account, error)
}

// constructors below

func NewAuthRepository(queries *model.Queries) *AuthRepository {
	return &AuthRepository{
		queries: queries,
	}
}
