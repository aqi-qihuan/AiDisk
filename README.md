# AiDisk

AI-powered cloud drive platform with intelligent agent capabilities.

## Project Structure

```
AiDisk/
├── AqiCloud-Agent/    # Agent management backend (Go + Gin)
├── AqiCloud-AI/       # AI agent backend (Go + LLM integration)
├── AqiCloud-Web/      # Web UI (Vue 3 + Vite)
├── bin/               # Build output directory
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
cd AqiCloud-Agent && go run ./cmd/server
cd AqiCloud-AI && go run ./cmd/agent
cd AqiCloud-Web && npm run dev
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
