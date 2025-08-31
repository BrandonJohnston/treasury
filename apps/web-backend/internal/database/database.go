package database

import (
	"database/sql"
	"fmt"
	"time"

	"web-backend/config"

	_ "github.com/lib/pq"
)

func SetupDB() (*sql.DB, error) {
	// Connection string (adjust user, password, dbname, host, port)
	// connStr := "postgres://user:password@localhost:5432/dbname?sslmode=disable"

	cfg := config.Load()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode)

	// Open the DB connection (does NOT establish network connection yet)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Ping to verify that the connection is actually alive
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Optional: Set connection pool options
	db.SetMaxOpenConns(25) // max open connections
	db.SetMaxIdleConns(25) // max idle connections
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

// var DB *sql.DB

// // Connect establishes a connection to the PostgreSQL database
// func Connect(config *config.Config) error {
// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.SSLMode)

// 	var err error
// 	DB, err = sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		return fmt.Errorf("error opening database: %v", err)
// 	}

// 	// Test the connection
// 	if err = DB.Ping(); err != nil {
// 		return fmt.Errorf("error connecting to the database: %v", err)
// 	}

// 	log.Println("Successfully connected to PostgreSQL database")
// 	return nil
// }

// // Close closes the database connection
// func Close() error {
// 	if DB != nil {
// 		return DB.Close()
// 	}
// 	return nil
// }
