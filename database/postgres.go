package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// PostgresConfig represents the configuration for a PostgreSQL database
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewPostgresDB returns a new PostgreSQL database connection
func NewPostgresDB(cfg PostgresConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	return sql.Open("postgres", connStr)
}

// PostgresDB is a wrapper around the sql.DB object for PostgreSQL
type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDBConnection returns a new PostgresDB object
func NewPostgresDBConnection(db *sql.DB) *PostgresDB {
	return &PostgresDB{db: db}
}

// Ping checks the connection to the database
func (p *PostgresDB) Ping(ctx context.Context) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return p.db.PingContext(ctxTimeout)
}

// ExecuteQuery executes a SQL query on the database
func (p *PostgresDB) ExecuteQuery(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return p.db.QueryContext(ctxTimeout, query, args...)
}

// Close closes the database connection
func (p *PostgresDB) Close() error {
	return p.db.Close()
}

func main() {}
