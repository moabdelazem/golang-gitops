package services

import (
	"database/sql"
	"fmt"

	"github.com/moabdelazem/golang-gitops/pkg/database"
)

// CounterService handles counter-related operations
type CounterService struct {
	db *database.DB
}

// NewCounterService creates a new counter service
func NewCounterService(db *database.DB) *CounterService {
	return &CounterService{db: db}
}

// Counter represents a counter record
type Counter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// IncrementCounter increments a counter by name and returns the new value
func (s *CounterService) IncrementCounter(name string) (*Counter, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection not available")
	}

	query := `
		UPDATE counters 
		SET value = value + 1, updated_at = CURRENT_TIMESTAMP 
		WHERE name = $1 
		RETURNING id, name, value
	`

	var counter Counter
	err := s.db.QueryRow(query, name).Scan(&counter.ID, &counter.Name, &counter.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			// If counter doesn't exist, create it
			return s.createCounter(name, 1)
		}
		return nil, fmt.Errorf("failed to increment counter: %w", err)
	}

	return &counter, nil
}

// GetCounter retrieves a counter by name
func (s *CounterService) GetCounter(name string) (*Counter, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection not available")
	}

	query := `SELECT id, name, value FROM counters WHERE name = $1`

	var counter Counter
	err := s.db.QueryRow(query, name).Scan(&counter.ID, &counter.Name, &counter.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return s.createCounter(name, 0)
		}
		return nil, fmt.Errorf("failed to get counter: %w", err)
	}

	return &counter, nil
}

// createCounter creates a new counter with the specified value
func (s *CounterService) createCounter(name string, value int) (*Counter, error) {
	query := `
		INSERT INTO counters (name, value) 
		VALUES ($1, $2) 
		RETURNING id, name, value
	`

	var counter Counter
	err := s.db.QueryRow(query, name, value).Scan(&counter.ID, &counter.Name, &counter.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to create counter: %w", err)
	}

	return &counter, nil
}
