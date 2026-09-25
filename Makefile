.PHONY: build-all build-agent build-ai build-frontend clean deps

BIN_DIR := bin

build-all: build-agent build-ai build-frontend

build-agent:
	@echo "Building AgentPan backend..."
	cd AqiCloud-Agent && go build -o ../$(BIN_DIR)/agentpan-server ./cmd/server

build-ai:
	@echo "Building AI backend..."
	cd AqiCloud-AI && go build -o ../$(BIN_DIR)/ai-agent ./cmd/agent

build-frontend:
	@echo "Building frontend..."
	cd AqiCloud-Web && npm run build

deps:
	@echo "Installing dependencies..."
	cd AqiCloud-Agent && go mod tidy
	cd AqiCloud-AI && go mod tidy
	cd AqiCloud-Web && npm install

clean:
	rm -rf $(BIN_DIR)/*
	cd frontend && rm -rf dist
