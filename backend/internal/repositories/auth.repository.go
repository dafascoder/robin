package repositories

import (
	model "backend/internal/models"
	"context"
	"github.com/google/uuid"
)

type AuthRepository struct {
	queries *model.Queries
}

func (ar *AuthRepository) CreateAccount(ctx context.Context, params model.CreateAccountParams) (model.CreateAccountRow, error) {
	account, err := ar.queries.CreateAccount(ctx, params)
	if err != nil {
		return model.CreateAccountRow{}, err
	}

	return account, nil
}

func (ar *AuthRepository) GetAccountByEmail(ctx context.Context, email string) (model.Account, error) {
	data, err := ar.queries.GetAccountByEmail(ctx, email)
	if err != nil {
		return model.Account{}, err
	}

	return data, nil
}

func (ar *AuthRepository) GetAccountByID(ctx context.Context, id uuid.UUID) (model.Account, error) {
	return model.Account{}, nil
}

func (ar *AuthRepository) VerifyAccount(ctx context.Context, accountID uuid.UUID) error {
	_, err := ar.queries.VerifyAccount(ctx, accountID)
	if err != nil {
		return err
	}

	return nil
}

func (ar *AuthRepository) UpdateRefreshToken(ctx context.Context, accountID uuid.UUID) error {
	_, err := ar.queries.UpdateRefreshTokenVersion(ctx, accountID)
	if err != nil {
		return err
	}

	return nil
}
