package database

import (
	"backend/internal/config"
	logging "backend/internal/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	db *pgxpool.Pool
}

var dbInstance *Database

func NewDatabase(ctx context.Context) (*Database, error) {
	db, err := pgxpool.New(ctx, config.Env.DATABASEURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	logging.Logger.LogDebug().Msg("Connected to database")

	dbInstance = &Database{
		db: db,
	}

	return dbInstance, nil
}

func (pg *Database) GetDatabaseInstance() *pgxpool.Pool {
	return dbInstance.db
}

func (pg *Database) Ping(ctx context.Context) error {
	logging.Logger.LogDebug().Msg("Pinging database")
	return pg.db.Ping(ctx)
}

func (pg *Database) Close() {
	pg.db.Close()
}
