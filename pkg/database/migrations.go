package database

import (
	"log"
)

// CreateTables creates the necessary tables for the application
func (db *DB) CreateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS counters (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL,
		value INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- Insert default counter if it doesn't exist
	INSERT INTO counters (name, value) 
	VALUES ('api_hits', 0) 
	ON CONFLICT (name) DO NOTHING;
	`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	log.Println("Database tables created successfully")
	return nil
}
