.PHONY: build-all build-agent build-ai build-frontend clean deps

BIN_DIR := bin

build-all: build-agent build-ai build-frontend

build-agent:
	@echo "Building AgentPan backend..."
	cd backend-agent && go build -o ../$(BIN_DIR)/agentpan-server ./cmd/server

build-ai:
	@echo "Building AI backend..."
	cd backend-ai && go build -o ../$(BIN_DIR)/ai-agent ./cmd/agent

build-frontend:
	@echo "Building frontend..."
	cd frontend && npm run build

deps:
	@echo "Installing dependencies..."
	cd backend-agent && go mod tidy
	cd backend-ai && go mod tidy
	cd frontend && npm install

clean:
	rm -rf $(BIN_DIR)/*
	cd frontend && rm -rf dist
