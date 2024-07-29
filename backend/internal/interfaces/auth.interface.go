package interfaces

import (
	model "backend/internal/models"
	"context"
)

type AuthInterface interface {
	CreateAccount(ctx context.Context, params model.CreateAccountParams) (model.CreateAccountRow, error)
	GetAccountByEmail(ctx context.Context, email string) (model.Account, error)
	GetAccountByID(ctx context.Context) (model.Account, error)
}
