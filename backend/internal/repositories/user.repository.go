package repositories

import (
	model "backend/internal/models"
	"context"
	"github.com/google/uuid"
)

type UserRepository struct {
	queries *model.Queries
}

func (ur *UserRepository) GetUserByAccountID(ctx context.Context, accountId uuid.UUID) (model.User, error) {
	user, err := ur.queries.GetUserByAccountID(ctx, accountId)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	user, err := ur.queries.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, user model.CreateUserParams) (model.CreateUserRow, error) {
	newUser, err := ur.queries.CreateUser(ctx, model.CreateUserParams{
		Name:      user.Name,
		AccountID: user.AccountID,
	})
	if err != nil {
		return model.CreateUserRow{}, err
	}

	return newUser, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, data model.UpdateUserByIDParams) error {
	err := ur.queries.UpdateUserByID(ctx, model.UpdateUserByIDParams{
		ID:   data.ID,
		Name: data.Name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := ur.queries.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
