package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/moabdelazem/golang-gitops/internal/services"
)

// CounterHandler handles counter-related HTTP requests
type CounterHandler struct {
	counterService *services.CounterService
}

// NewCounterHandler creates a new counter handler
func NewCounterHandler(counterService *services.CounterService) *CounterHandler {
	return &CounterHandler{
		counterService: counterService,
	}
}

// CounterResponse represents the response for counter operations
type CounterResponse struct {
	Success bool              `json:"success"`
	Data    *services.Counter `json:"data,omitempty"`
	Error   string            `json:"error,omitempty"`
	Message string            `json:"message,omitempty"`
}

// HandleIncrement increments the API hits counter
func (h *CounterHandler) HandleIncrement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	counter, err := h.counterService.IncrementCounter("api_hits")
	if err != nil {
		log.Printf("Error incrementing counter: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CounterResponse{
			Success: false,
			Error:   "Failed to increment counter",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CounterResponse{
		Success: true,
		Data:    counter,
		Message: "Counter incremented successfully",
	})

	log.Printf("Counter incremented: %s = %d", counter.Name, counter.Value)
}

// HandleGetCounter retrieves the current counter value
func (h *CounterHandler) HandleGetCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	counter, err := h.counterService.GetCounter("api_hits")
	if err != nil {
		log.Printf("Error getting counter: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CounterResponse{
			Success: false,
			Error:   "Failed to get counter",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CounterResponse{
		Success: true,
		Data:    counter,
		Message: "Counter retrieved successfully",
	})
}
