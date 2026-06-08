package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config represents the application configuration
.type Config struct {
	PostgresHost string `json:"postgres_host"`
	PostgresPort string `json:"postgres_port"`
	PostgresUser string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
	PostgresDatabase string `json:"postgres_database"`
	SQLiteDatabase string `json:"sqlite_database"`
	RedisHost string `json:"redis_host"`
	RedisPort string `json:"redis_port"`
	RedisPassword string `json:"redis_password"`
	PrismaSchema string `json:"prisma_schema"`
}

// LoadConfig loads the application configuration from a JSON file
func LoadConfig(filename string) (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}

// SaveConfig saves the application configuration to a JSON file
func SaveConfig(config *Config, filename string) error {
	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// GetConfig returns the application configuration
func GetConfig() *Config {
	configFile := "config.json"
	config, err := LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	// Example usage
	config := GetConfig()
	fmt.Printf("Postgres Host: %s\n", config.PostgresHost)
	fmt.Printf("Postgres Port: %s\n", config.PostgresPort)
	fmt.Printf("Postgres User: %s\n", config.PostgresUser)
	fmt.Printf("Postgres Password: %s\n", config.PostgresPassword)
	fmt.Printf("Postgres Database: %s\n", config.PostgresDatabase)
	fmt.Printf("SQLite Database: %s\n", config.SQLiteDatabase)
	fmt.Printf("Redis Host: %s\n", config.RedisHost)
	fmt.Printf("Redis Port: %s\n", config.RedisPort)
	fmt.Printf("Redis Password: %s\n", config.RedisPassword)
	fmt.Printf("Prisma Schema: %s\n", config.PrismaSchema)
}