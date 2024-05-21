package database

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", "postgres://irfanmuhammadfauzi@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		slog.Error("Failed to connect DB", "Error", err)
		return db, err
	}
	return db, nil
}
