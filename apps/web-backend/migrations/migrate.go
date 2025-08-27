package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "password"
	dbname   = "treasury"
	sslmode  = "disable"
)

// Migration represents a database migration
type Migration struct {
	ID       string
	Filename string
	SQL      string
}

func main() {
	// Connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// Create migrations table if it doesn't exist
	if err := createMigrationsTable(db); err != nil {
		log.Fatalf("Error creating migrations table: %v", err)
	}

	// Get all migration files
	migrations, err := getMigrationFiles()
	if err != nil {
		log.Fatalf("Error reading migration files: %v", err)
	}

	// Get applied migrations
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		log.Fatalf("Error getting applied migrations: %v", err)
	}

	// Apply pending migrations
	for _, migration := range migrations {
		if !isMigrationApplied(appliedMigrations, migration.ID) {
			log.Printf("Applying migration: %s", migration.Filename)
			if err := applyMigration(db, migration); err != nil {
				log.Fatalf("Error applying migration %s: %v", migration.Filename, err)
			}
			log.Printf("✅ Successfully applied migration: %s", migration.Filename)
		} else {
			log.Printf("⏭️  Migration already applied: %s", migration.Filename)
		}
	}

	log.Println("🎉 All migrations completed successfully!")
}

func createMigrationsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS migrations (
		id VARCHAR(255) PRIMARY KEY,
		filename VARCHAR(255) NOT NULL,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	return err
}

func getMigrationFiles() ([]Migration, error) {
	migrationsDir := "migrations"
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	var migrations []Migration
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			content, err := ioutil.ReadFile(filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				return nil, err
			}

			// Extract migration ID from filename (e.g., "001_create_users_table.sql" -> "001")
			parts := strings.Split(file.Name(), "_")
			if len(parts) > 0 {
				migrationID := parts[0]
				migrations = append(migrations, Migration{
					ID:       migrationID,
					Filename: file.Name(),
					SQL:      string(content),
				})
			}
		}
	}

	// Sort migrations by ID
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].ID < migrations[j].ID
	})

	return migrations, nil
}

func getAppliedMigrations(db *sql.DB) ([]string, error) {
	query := "SELECT id FROM migrations ORDER BY id"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applied []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		applied = append(applied, id)
	}

	return applied, nil
}

func isMigrationApplied(appliedMigrations []string, migrationID string) bool {
	for _, applied := range appliedMigrations {
		if applied == migrationID {
			return true
		}
	}
	return false
}

func applyMigration(db *sql.DB, migration Migration) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Execute the migration SQL
	if _, err := tx.Exec(migration.SQL); err != nil {
		tx.Rollback()
		return err
	}

	// Record the migration as applied
	recordQuery := "INSERT INTO migrations (id, filename) VALUES ($1, $2)"
	if _, err := tx.Exec(recordQuery, migration.ID, migration.Filename); err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit()
}
