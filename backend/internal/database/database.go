package database

import (
	"backend/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	db *sql.DB
}

func New() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Env.DBHost, config.Env.DBPort, config.Env.DBUser, config.Env.DBPass, config.Env.DBName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	return db, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (d *Database) Close() error {
	log.Printf("Disconnected from database: %s", config.Env.DBName)
	return d.Close()
}
