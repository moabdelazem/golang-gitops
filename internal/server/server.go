package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/moabdelazem/golang-gitops/internal/handlers"
	"github.com/moabdelazem/golang-gitops/pkg/database"
)

// Server holds the HTTP server and its dependencies
type Server struct {
	router *mux.Router
	db     *database.DB
	config Config
}

// Config holds server configuration
type Config struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// New creates a new server instance
func New(cfg Config, db *database.DB) *Server {
	s := &Server{
		router: mux.NewRouter(),
		db:     db,
		config: cfg,
	}

	s.setupRoutes()
	return s
}

// setupRoutes configures all the routes
func (s *Server) setupRoutes() {
	healthHandler := handlers.NewHealthHandler(s.db)

	s.router.HandleFunc("/health", healthHandler.HandleHealthCheck).Methods("GET")
	s.router.HandleFunc("/healthz", healthHandler.HandleHealthCheck).Methods("GET") // Kubernetes style
}

// Start starts the HTTP server
func (s *Server) Start() error {
	srv := &http.Server{
		Addr:         s.config.Address,
		Handler:      s.router,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
		IdleTimeout:  s.config.IdleTimeout,
	}

	// Channel to listen for interrupt signal to gracefully shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s", s.config.Address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Println("Server started successfully")

	// Wait for interrupt signal
	<-done
	log.Println("Server is shutting down...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server exited gracefully")
	return nil
}

// DefaultConfig returns a default server configuration
func DefaultConfig(address string) Config {
	return Config{
		Address:      address,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
