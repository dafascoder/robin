package services

import (
	model "backend/internal/models"
	"backend/internal/repositories"
	"context"
	"github.com/google/uuid"
)

type UserServices struct {
	repo *repositories.UserRepository
}

func (us *UserServices) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	user, err := us.repo.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *UserServices) GetUserByAccountID(ctx context.Context, accountId uuid.UUID) (model.User, error) {
	user, err := us.repo.GetUserByAccountID(ctx, accountId)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *UserServices) CreateUser(ctx context.Context, user model.CreateUserParams) (model.CreateUserRow, error) {
	newUser, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return model.CreateUserRow{}, err
	}

	return newUser, nil
}

func (us *UserServices) UpdateUser(ctx context.Context, data model.UpdateUserByIDParams) error {
	err := us.repo.UpdateUser(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserServices) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := us.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
