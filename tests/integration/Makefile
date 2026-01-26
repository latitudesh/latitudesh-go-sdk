# Makefile for Integration Tests

.PHONY: help test test-verbose test-coverage test-plans test-apikeys test-projects test-teams test-userdata test-servers clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run all tests
	@echo "Running all integration tests..."
	@go test -v ./...

test-verbose: ## Run all tests with verbose output
	@echo "Running all integration tests (verbose)..."
	@go test -v -count=1 ./...

test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	@go test -v -cover ./...
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-plans: ## Run Plans tests only
	@echo "Running Plans integration tests..."
	@go test -v -run TestPlansIntegration

test-apikeys: ## Run API Keys tests only
	@echo "Running API Keys integration tests..."
	@go test -v -run TestAPIKeysIntegration

test-projects: ## Run Projects tests only
	@echo "Running Projects integration tests..."
	@go test -v -run TestProjectsIntegration

test-teams: ## Run Teams tests only
	@echo "Running Teams integration tests..."
	@go test -v -run TestTeamsIntegration

test-userdata: ## Run User Data tests only
	@echo "Running User Data integration tests..."
	@go test -v -run TestUserDataIntegration

test-servers: ## Run Servers tests only
	@echo "Running Servers integration tests..."
	@go test -v -run TestServersIntegration

test-short: ## Run tests in short mode
	@echo "Running tests in short mode..."
	@go test -v -short ./...

test-race: ## Run tests with race detector
	@echo "Running tests with race detector..."
	@go test -v -race ./...

benchmark: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

clean: ## Clean test artifacts
	@echo "Cleaning test artifacts..."
	@rm -f coverage.out coverage.html
	@go clean -testcache

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@cd ../.. && go mod download
	@cd ../.. && go mod tidy

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run ./... || echo "golangci-lint not installed, skipping..."

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

check: fmt vet lint ## Run all code quality checks

all: check test ## Run all checks and tests
