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
	VerifyAccount(ctx context.Context, accountID uuid.UUID) error
}

type UserRepositoryInterface interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error)
	GetUserByAccountID(ctx context.Context, accountId uuid.UUID) (model.User, error)
	CreateUser(ctx context.Context, user model.CreateUserParams) (model.CreateUserRow, error)
	UpdateUser(ctx context.Context, data model.UpdateUserByIDParams) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// constructors below

func NewAuthRepository(queries *model.Queries) *AuthRepository {
	return &AuthRepository{
		queries: queries,
	}
}

func NewUserRepository(queries *model.Queries) *UserRepository {
	return &UserRepository{
		queries: queries,
	}
}
