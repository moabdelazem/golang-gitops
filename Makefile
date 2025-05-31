# Variables
BINARY_NAME=main
BUILD_DIR=bin

# Default target
.DEFAULT_GOAL := run

run:
	@go run cmd/main.go

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/main.go
	@echo "Built: $(BUILD_DIR)/$(BINARY_NAME)"

start: build
	@./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned build directory"

test:
	@go test ./...

fmt:
	@go fmt ./...

deps:
	@go mod download
	@go mod tidy

dev-setup:
	@cp .env.example .env
	@echo "Development environment setup complete. Please update .env with your database credentials."

.PHONY: run build start clean test fmt deps dev-setup