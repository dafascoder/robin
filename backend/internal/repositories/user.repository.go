package repositories

import (
	model "backend/internal/models"
	"context"
	"github.com/google/uuid"
)

type UserRepository struct {
	queries *model.Queries
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	user, err := ur.queries.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
