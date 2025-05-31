package main

import (
	"log"

	"github.com/moabdelazem/golang-gitops/internal/server"
	"github.com/moabdelazem/golang-gitops/pkg/config"
	"github.com/moabdelazem/golang-gitops/pkg/database"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database connection
	dbConfig := database.DefaultConfig(cfg.GetDSN())
	db, err := database.New(dbConfig)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		log.Println("Starting server without database connection")
		db = nil
	}

	// Ensure database connection is closed when main exits
	if db != nil {
		defer db.Close()
	}

	// Create and start server
	serverConfig := server.DefaultConfig(cfg.GetServerAddress())
	srv := server.New(serverConfig, db)

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
