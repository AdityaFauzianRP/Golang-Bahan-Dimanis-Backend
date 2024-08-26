package config

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// Config holds the configuration values
type Config struct {
	DatabaseURL string `json:"database_url"`
}

var DB *sql.DB

// LoadConfig loads configuration from a JSON file
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

// ConnectDB initializes and returns a connection to the PostgreSQL database
func ConnectDB(config *Config) (*sql.DB, error) {
	connStr := config.DatabaseURL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Ensure the database connection is available
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	DB = db
	return db, nil
}
