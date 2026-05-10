#!/usr/bin/env bash
set -euo pipefail

# AiDisk deploy script
# Builds all services, packages them into a tarball, and optionally uploads.

BIN_DIR=bin
OUTPUT="deploy-$(date +%Y%m%d%H%M).tar.gz"

echo "=== AiDisk Build & Package ==="

# Build
make build-all

# Package
echo "Packaging: $OUTPUT"
tar czf "$OUTPUT" \
  "$BIN_DIR"/agentpan-server \
  "$BIN_DIR"/ai-agent \
  frontend/dist/ \
  docker-compose.yml \
  .env.example \
  nginx*.conf 2>/dev/null || true

echo "Done: $OUTPUT"
echo ""
echo "Upload to server:"
echo "  scp $OUTPUT user@server:/opt/aidisk/"
echo ""
echo "On server, extract:"
echo "  tar xzf $OUTPUT -C /opt/aidisk/"
