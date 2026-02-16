package internal

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Connect establishes a connection to the PostgreSQL database.
func Connect(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

// RunMigrations executes database migrations.
func RunMigrations(db *sql.DB) error {
	// Implement migration logic here.
	return nil
}