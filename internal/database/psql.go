package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

func NewPostgreSQLConnection(dbSourse string) *sql.DB {
	slog.Info("Connecting to PostgreSQL database")

	db, err := sql.Open("postgres", dbSourse)
	if err != nil {
		panic(fmt.Sprintf("Failed to open PostgreSQL connection: %v", err))
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping PostgreSQL database: %v", err))
	}

	slog.Info("Successfully connected to PostgreSQL database")
	return db
}
