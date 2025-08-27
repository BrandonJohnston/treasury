package database

import (
	"database/sql"
	"fmt"
	"log"

	"web-backend/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Connect establishes a connection to the PostgreSQL database
func Connect(config *config.Config) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.SSLMode)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
