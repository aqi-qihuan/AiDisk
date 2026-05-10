# AiDisk

AI-powered cloud drive platform with intelligent agent capabilities.

## Project Structure

```
AiDisk/
├── backend-agent/       # Agent management backend (Go + Gin)
├── backend-ai/          # AI agent backend (Go + LLM integration)
├── frontend/            # Web UI (Vue 3 + Vite)
├── bin/                 # Build output directory
├── Makefile             # Root build commands
└── deploy.sh            # Build & package script
```

## Quick Start

### Prerequisites

- Go >= 1.24
- Node.js >= 18
- MySQL, Redis, MinIO (for running services)

### Build

```bash
# Install dependencies
make deps

# Build all services
make build-all

# Build individually
make build-agent    # AgentPan backend
make build-ai       # AI backend
make build-frontend # Web UI
```

### Local Development

Each service can be run independently:

```bash
cd backend-agent && go run ./cmd/server
cd backend-ai && go run ./cmd/agent
cd frontend && npm run dev
```

### Deployment

```bash
# Package all services
./deploy.sh

# Upload to server
scp deploy-*.tar.gz user@server:/opt/aidisk/

# Extract on server
tar xzf deploy-*.tar.gz -C /opt/aidisk/
```

See `.env.example` for required environment variables.

## License

MIT
