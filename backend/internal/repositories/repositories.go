package repositories

import (
	model "backend/internal/models"
)

func NewAuthRepository(queries *model.Queries) *AuthRepository {
	return &AuthRepository{
		queries: queries,
	}
}
