package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/moabdelazem/golang-gitops/pkg/database"
)

// HealthHandler holds dependencies for health check handler
type HealthHandler struct {
	db *database.DB
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(db *database.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
	Message  string `json:"message,omitempty"`
}

// HandleHealthCheck responds to health check requests
func (h *HealthHandler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status: "OK",
	}

	// Check database health
	if h.db != nil {
		if err := h.db.HealthCheck(); err != nil {
			response.Database = "DOWN"
			response.Message = "Database connection failed"
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			response.Database = "UP"
		}
	} else {
		response.Database = "NOT_CONFIGURED"
	}

	if response.Database == "UP" {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
}
