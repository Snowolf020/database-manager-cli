package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteDatabase represents a SQLite database connection
// and provides methods for interacting with it.
type SQLiteDatabase struct {
	db *sql.DB
}

// NewSQLiteDatabase returns a new SQLiteDatabase instance.
func NewSQLiteDatabase(dbPath string) (*SQLiteDatabase, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &SQLiteDatabase{db: db}, nil
}

// Ping checks if the database connection is alive.
func (s *SQLiteDatabase) Ping() error {
	return s.db.Ping()
}

// ExecuteQuery executes a SQL query on the database.
func (s *SQLiteDatabase) ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}

// ExecuteUpdate executes a SQL update statement on the database.
func (s *SQLiteDatabase) ExecuteUpdate(query string, args ...interface{}) (int64, error) {
	result, err := s.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Close closes the database connection.
func (s *SQLiteDatabase) Close() error {
	return s.db.Close()
}

func main() {
	// Create a new SQLite database connection
	db, err := NewSQLiteDatabase("example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to check if it's alive
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Execute a query on the database
	rows, err := db.ExecuteQuery("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the results of the query
	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}