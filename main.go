package main

import (
	"database-manager-cli/config"
	"database-manager-cli/database"
	"flag"
	"fmt"
	"log"
)

func main() {
	configPath := flag.String("config", "config/config.json", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	switch cfg.Database.Type {
	case "postgres":
		db, err := database.NewPostgresDB(cfg.Database.Postgres)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	case "sqlite":
		db, err := database.NewSqliteDB(cfg.Database.Sqlite)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	default:
		log.Fatal("Unsupported database type")
	}

	// Start the command line tool
	cmd := database.NewCLI(db)
	cmd.Start()
}
