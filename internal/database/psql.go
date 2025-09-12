package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConnection(source string) *sql.DB {
	slog.Info("Connecting to SQLite database (sqlx)")

	db, err := sql.Open("sqlite3", source)
	if err != nil {
		panic(err)
	}

	// optional pragmas
	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	slog.Info("Successfully connected to SQLite database")
	return db
}

func NewEmbeddedSQLite(b []byte) (*sql.DB, error) {

	tmpFile, err := os.CreateTemp("", "app-*.db")
	if err != nil {
		return nil, err
	}

	if _, err := tmpFile.Write(b); err != nil {
		return nil, err
	}
	tmpFile.Close()

	db, err := sql.Open("sqlite3", tmpFile.Name())
	if err != nil {
		return nil, err
	}

	return db, nil
}
