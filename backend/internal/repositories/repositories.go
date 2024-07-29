package repositories

import (
	model "backend/internal/models"
	"database/sql"
)

func NewAuthRepository(db *sql.DB, queries *model.Queries) *AuthRepository {
	return &AuthRepository{
		db:      db,
		queries: queries,
	}
}
