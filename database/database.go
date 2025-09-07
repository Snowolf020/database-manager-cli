package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// DB is a database connection wrapper
type DB struct {
	db *sqlx.DB
}

// NewDB returns a new database connection
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	// Test the database connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

// NewSQLiteDB returns a new SQLite database connection
func NewSQLiteDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	// Test the database connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

// Query executes a SQL query and returns the results
func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	rows, err := d.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Exec executes a SQL query and returns the result
func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	result, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Close closes the database connection
func (d *DB) Close() error {
	return d.db.Close()
}

// GetPostgresDB returns a Postgres database connection
func GetPostgresDB(dataSourceName string) (*DB, error) {
	return NewDB(dataSourceName)
}

// GetSQLiteDB returns a SQLite database connection
func GetSQLiteDB(dataSourceName string) (*DB, error) {
	return NewSQLiteDB(dataSourceName)
}
