package schema

import (
	"database/sql"
	"errors"
	"log"
	"path/filepath"

	"github.com/prisma/prisma-client-go/prisma"
)

// SchemaManager handles database schema management logic
	type SchemaManager struct {
		db        *sql.DB
		schemaPath string
	}

// NewSchemaManager returns a new instance of SchemaManager
func NewSchemaManager(db *sql.DB, schemaPath string) *SchemaManager {
	return &SchemaManager{db: db, schemaPath: schemaPath}
}

// CreateSchema creates the database schema
func (sm *SchemaManager) CreateSchema() error {
	// Load schema from file
		schema, err := filepath.Glob(sm.schemaPath + "/*.prisma")
	if err != nil {
		return err
	}

	// Create Prisma client
		prismaClient, err := prisma.New(sm.db, schema[0])
	if err != nil {
		return err
	}

	// Create schema
		_, err = prismaClient.PrismaMigrate("init")
	if err != nil {
		return err
	}

	return nil
}

// DropSchema drops the database schema
func (sm *SchemaManager) DropSchema() error {
	// Create Prisma client
		prismaClient, err := prisma.New(sm.db, sm.schemaPath+"/schema.prisma")
	if err != nil {
		return err
	}

	// Drop schema
		_, err = prismaClient.PrismaMigrate("reset")
	if err != nil {
		return err
	}

	return nil
}